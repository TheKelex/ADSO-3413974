// Portable Docker Compose validation runner for the database project.
// Starts PostgreSQL, applies forward migrations and the seed, then runs
// schema assertions. Uses only Node stdlib and the Docker CLI; it imports
// no external database client. Run from the database/ directory.
import { execFileSync } from 'node:child_process';
import { readFileSync, readdirSync } from 'node:fs';
import { join } from 'node:path';
import { pathToFileURL } from 'node:url';

const ENV_FILE = '.env.example';
const COMPOSE = ['compose', '--env-file', ENV_FILE, '-f', 'docker-compose.db.yml'];

function loadEnv(file) {
  const env = {};
  for (const line of readFileSync(file, 'utf8').split(/\r?\n/)) {
    const text = line.trim();
    if (!text || text.startsWith('#')) continue;
    const eq = text.indexOf('=');
    if (eq === -1) continue;
    env[text.slice(0, eq).trim()] = text.slice(eq + 1).trim();
  }
  return env;
}

const env = loadEnv(ENV_FILE);
function requiredEnv(name) {
  const value = env[name];
  if (!value) throw new Error(`Missing required variable ${name} in ${ENV_FILE}`);
  return value;
}

function docker(args, options = {}) {
  return execFileSync('docker', [...COMPOSE, ...args], { encoding: 'utf8', ...options });
}

function runSql(sqlText, vars = {}) {
  const varArgs = [];
  for (const [key, value] of Object.entries(vars)) varArgs.push('-v', `${key}=${value}`);
  return execFileSync(
    'docker',
    [...COMPOSE, 'exec', '-T', 'db', 'psql', '-v', 'ON_ERROR_STOP=1', ...varArgs,
      '-U', requiredEnv('POSTGRES_USER'), '-d', requiredEnv('POSTGRES_DB')],
    { input: sqlText, encoding: 'utf8' },
  );
}

function queryScalar(sql) {
  return execFileSync(
    'docker',
    [...COMPOSE, 'exec', '-T', 'db', 'psql', '-tA',
      '-U', requiredEnv('POSTGRES_USER'), '-d', requiredEnv('POSTGRES_DB'), '-c', sql],
    { encoding: 'utf8' },
  ).trim();
}

const wait = (ms) => new Promise((resolve) => setTimeout(resolve, ms));

async function waitForDatabase() {
  for (let attempt = 0; attempt < 30; attempt += 1) {
    try {
      execFileSync(
        'docker',
        [...COMPOSE, 'exec', '-T', 'db', 'pg_isready',
          '-U', requiredEnv('POSTGRES_USER'), '-d', requiredEnv('POSTGRES_DB')],
        { stdio: 'ignore' },
      );
      return;
    } catch {
      await wait(2000);
    }
  }
  throw new Error('PostgreSQL did not become ready within the timeout.');
}

function applyMigrations() {
  const files = readdirSync('migrations').filter((f) => f.endsWith('.up.sql')).sort();
  for (const file of files) runSql(readFileSync(join('migrations', file), 'utf8'));
  return files.length;
}

function applySeed() {
  runSql(readFileSync(join('seed', '0001_seed_admin.sql'), 'utf8'), {
    admin_username: requiredEnv('ADMIN_USERNAME'),
    admin_password: requiredEnv('ADMIN_PASSWORD'),
  });
}

async function main() {
  docker(['up', '-d']);
  try {
    await waitForDatabase();
    const applied = applyMigrations();
    applySeed();
    const url = pathToFileURL(join(process.cwd(), 'tests', 'schema_assertions.mjs')).href;
    const { runAssertions } = await import(url);
    runAssertions(queryScalar);
    console.log(`Validation passed: ${applied} migrations applied, seed and schema verified.`);
  } finally {
    try { docker(['down', '-v']); } catch { /* best-effort cleanup */ }
  }
}

main().catch((error) => {
  console.error(error.message || error);
  process.exit(1);
});
