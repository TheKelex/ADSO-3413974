// Assertions invoked by scripts/run_validation.mjs after migrations and seed.
// queryScalar(sql) runs psql -tA inside the container and returns the value.
export function runAssertions(queryScalar) {
  const expected = [
    'user', 'passenger', 'airport', 'aircraft', 'reservation', 'ticket',
    'flight', 'seat', 'seat_assignment', 'payment', 'baggage', 'boarding',
  ];
  const list = expected.map((t) => `'${t}'`).join(', ');
  const tableCount = queryScalar(
    "SELECT count(*) FROM information_schema.tables " +
    `WHERE table_schema = 'public' AND table_name IN (${list})`,
  );
  if (tableCount !== String(expected.length)) {
    throw new Error(`Expected ${expected.length} tables, found ${tableCount}`);
  }

  const adminCount = queryScalar("SELECT count(*) FROM \"user\" WHERE role = 'admin'");
  if (Number(adminCount) < 1) {
    throw new Error('Expected at least one seeded admin user, found none');
  }

  const bcryptCount = queryScalar(
    "SELECT count(*) FROM \"user\" WHERE role = 'admin' AND password_hash LIKE '$2%'",
  );
  if (Number(bcryptCount) < 1) {
    throw new Error('Seeded admin password is not stored as a bcrypt hash');
  }

  // Composite key sanity: flight PK spans two columns.
  const flightPkCols = queryScalar(
    "SELECT count(*) FROM information_schema.key_column_usage " +
    "WHERE table_name = 'flight' AND constraint_name LIKE '%pkey%'",
  );
  if (Number(flightPkCols) < 2) {
    throw new Error(`Expected a composite primary key on flight, found ${flightPkCols} column(s)`);
  }

  console.log('Schema assertions passed: 12 tables, composite flight key, admin seeded with bcrypt.');
}
