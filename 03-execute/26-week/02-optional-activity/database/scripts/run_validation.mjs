// Portable Docker Compose validation runner for the database project.
// Starts MySQL, waits for a real readiness signal, applies the forward
// migrations and the seed, then runs schema assertions against the live
// database. It uses only the Node standard library and the Docker CLI and
// imports no external database client. Run it from the database/ directory:
//   node scripts/run_validation.mjs
import { execFileSync } from 'node:child_process';
import { readFileSync, readdirSync } from 'node:fs';
import { join } from 'node:path';
import { pathToFileURL } from 'node:url';

const ENV_FILE = '.env.example';
const COMPOSE_FILE = 'docker-compose.db.yml';
const SERVICE = 'db';
const COMPOSE = ['compose', '--env-file', ENV_FILE, '-f', COMPOSE_FILE];
const SINGLE_QUOTE = String.fromCharCode(39);
const DOUBLE_QUOTE = String.fromCharCode(34);
const BACKSLASH = String.fromCharCode(92);

function loadEnvFile(file) {
  const values = {};
  for (const line of readFileSync(file, 'utf8').split(/\r?\n/)) {
    const text = line.trim();
    if (!text || text.startsWith('#')) continue;
    const separator = text.indexOf('=');
    if (separator === -1) continue;
    const key = text.slice(0, separator).trim();
    let value = text.slice(separator + 1).trim();
    const quoted =
      value.length > 1 &&
      ((value.startsWith(SINGLE_QUOTE) && value.endsWith(SINGLE_QUOTE)) ||
        (value.startsWith(DOUBLE_QUOTE) && value.endsWith(DOUBLE_QUOTE)));
    if (quoted) {
      value = value.slice(1, -1);
    }
    values[key] = value;
  }
  return values;
}

const env = loadEnvFile(ENV_FILE);

function requiredEnv(name) {
  const value = env[name];
  if (!value) {
    throw new Error('Missing required variable ' + name + ' in ' + ENV_FILE);
  }
  return value;
}

function compose(args, options = {}) {
  return execFileSync('docker', [...COMPOSE, ...args], { encoding: 'utf8', ...options });
}

function mysqlArgs(extra) {
  return [
    ...COMPOSE,
    'exec',
    '-T',
    '-e',
    'MYSQL_PWD=' + requiredEnv('MYSQL_PASSWORD'),
    SERVICE,
    'mysql',
    '-u',
    requiredEnv('MYSQL_USER'),
    '--default-character-set=utf8mb4',
    ...extra,
    requiredEnv('MYSQL_DATABASE'),
  ];
}

function runSql(sqlText) {
  return execFileSync('docker', mysqlArgs([]), { input: sqlText, encoding: 'utf8' });
}

function expectSqlError(sqlText) {
  try {
    execFileSync('docker', mysqlArgs([]), { input: sqlText, encoding: 'utf8', stdio: 'pipe' });
  } catch {
    return true;
  }
  return false;
}

function queryScalar(sql) {
  return execFileSync('docker', mysqlArgs(['-N', '-B', '-e', sql]), { encoding: 'utf8' }).trim();
}

const pause = (ms) => new Promise((resolve) => setTimeout(resolve, ms));

async function waitForDatabase() {
  const probe = [
    ...COMPOSE,
    'exec',
    '-T',
    '-e',
    'MYSQL_PWD=' + requiredEnv('MYSQL_ROOT_PASSWORD'),
    SERVICE,
    'mysqladmin',
    'ping',
    '-h',
    '127.0.0.1',
    '-P',
    '3306',
    '-u',
    'root',
    '--silent',
  ];
  for (let attempt = 0; attempt < 40; attempt += 1) {
    try {
      execFileSync('docker', probe, { stdio: 'ignore' });
      return;
    } catch {
      await pause(2000);
    }
  }
  const logs = compose(['logs', '--tail', '80', SERVICE]);
  throw new Error('MySQL did not become ready within the timeout.' + String.fromCharCode(10) + logs);
}

function applyMigrations() {
  const files = readdirSync('migrations')
    .filter((name) => name.endsWith('.up.sql'))
    .sort();
  for (const file of files) {
    runSql(readFileSync(join('migrations', file), 'utf8'));
  }
  return files.length;
}

function sqlLiteral(value) {
  const escaped = String(value)
    .split(BACKSLASH)
    .join(BACKSLASH + BACKSLASH)
    .split(SINGLE_QUOTE)
    .join(SINGLE_QUOTE + SINGLE_QUOTE);
  return SINGLE_QUOTE + escaped + SINGLE_QUOTE;
}

function applySeed() {
  const binding = [
    ['admin_username', requiredEnv('ADMIN_USERNAME')],
    ['admin_full_name', requiredEnv('ADMIN_FULL_NAME')],
    ['admin_password_hash', requiredEnv('ADMIN_PASSWORD_HASH')],
    ['technician_one_username', requiredEnv('TECHNICIAN_ONE_USERNAME')],
    ['technician_one_full_name', requiredEnv('TECHNICIAN_ONE_FULL_NAME')],
    ['technician_one_specialty', requiredEnv('TECHNICIAN_ONE_SPECIALTY')],
    ['technician_one_password_hash', requiredEnv('TECHNICIAN_ONE_PASSWORD_HASH')],
    ['technician_two_username', requiredEnv('TECHNICIAN_TWO_USERNAME')],
    ['technician_two_full_name', requiredEnv('TECHNICIAN_TWO_FULL_NAME')],
    ['technician_two_specialty', requiredEnv('TECHNICIAN_TWO_SPECIALTY')],
    ['technician_two_password_hash', requiredEnv('TECHNICIAN_TWO_PASSWORD_HASH')],
  ];
  const header = binding
    .map((entry) => 'SET @' + entry[0] + ' = ' + sqlLiteral(entry[1]) + ';')
    .join(String.fromCharCode(10));
  const seedFile = join('seed', '0001_seed_bootstrap.sql');
  runSql(header + String.fromCharCode(10) + readFileSync(seedFile, 'utf8'));
}

async function main() {
  try {
    compose(['up', '-d', '--wait']);
  } catch {
    compose(['up', '-d']);
  }
  try {
    await waitForDatabase();
    const applied = applyMigrations();
    applySeed();
    const assertionsUrl = pathToFileURL(join(process.cwd(), 'tests', 'schema_assertions.mjs')).href;
    const { runAssertions } = await import(assertionsUrl);
    const checks = runAssertions({ queryScalar, runSql, expectSqlError, env });
    console.log(
      'Validation passed: ' + applied + ' migrations applied, seed inserted, ' + checks + ' schema assertions verified.',
    );
  } finally {
    try {
      compose(['down', '-v']);
    } catch {
      // best effort cleanup, the validation result is already decided
    }
  }
}

main().catch((error) => {
  console.error(error.message || error);
  process.exit(1);
});
