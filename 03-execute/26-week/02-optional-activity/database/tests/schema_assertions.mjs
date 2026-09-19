// Schema assertions executed against the live MySQL instance after the forward
// migrations and the seed have been applied. They verify the structures the
// product depends on and the invariants the storage layer must enforce by
// itself: uniqueness, referential integrity, the lifecycle value set, the
// positive quantity rules and the one active assignment per technician rule.
const TABLE_NAME = [
  'user',
  'customer',
  'vehicle',
  'technician',
  'service_order',
  'assignment',
  'diagnostic',
  'intervention',
  'part_usage',
  'warranty',
  'status_transition',
];

const UNIQUE_INDEX = [
  ['user', 'uq_user_username'],
  ['customer', 'uq_customer_document_number'],
  ['vehicle', 'uq_vehicle_plate'],
  ['vehicle', 'uq_vehicle_vin'],
  ['technician', 'uq_technician_user_id'],
  ['service_order', 'uq_service_order_order_number'],
  ['assignment', 'uq_assignment_active_marker'],
];

const READ_INDEX = [
  ['service_order', 'ix_service_order_status'],
  ['service_order', 'ix_service_order_vehicle_id'],
  ['warranty', 'ix_warranty_expiration_date'],
  ['status_transition', 'ix_status_transition_service_order_id'],
];

const CUSTOMER_ID = 'c0000000-0000-4000-8000-000000000001';
const VEHICLE_ID = 'e0000000-0000-4000-8000-000000000001';
const ORDER_ONE_ID = '50000000-0000-4000-8000-000000000001';
const ORDER_TWO_ID = '50000000-0000-4000-8000-000000000002';
const ORDER_THREE_ID = '50000000-0000-4000-8000-000000000003';
const TECHNICIAN_ONE_ID = 'aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaa1';
const TECHNICIAN_TWO_ID = 'aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaa2';
const ADMIN_ID = '11111111-1111-4111-8111-111111111111';
const INTERVENTION_ID = '70000000-0000-4000-8000-000000000001';
const WARRANTY_ID = '80000000-0000-4000-8000-000000000001';

const USER_TABLE = String.fromCharCode(96) + 'user' + String.fromCharCode(96);

function assert(condition, message) {
  if (!condition) {
    throw new Error('Assertion failed: ' + message);
  }
}

export function runAssertions(context) {
  const { queryScalar, runSql, expectSqlError, env } = context;
  let checks = 0;

  for (const table of TABLE_NAME) {
    const found = queryScalar(
      "SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name = '" + table + "'",
    );
    assert(found === '1', 'table ' + table + ' exists');
    checks += 1;
  }

  for (const [table, index] of [...UNIQUE_INDEX, ...READ_INDEX]) {
    const found = queryScalar(
      "SELECT COUNT(*) FROM information_schema.statistics WHERE table_schema = DATABASE() AND table_name = '"
        + table + "' AND index_name = '" + index + "'",
    );
    assert(found !== '0', 'index ' + index + ' exists on ' + table);
    checks += 1;
  }

  const adminRole = queryScalar(
    'SELECT role FROM ' + USER_TABLE + " WHERE id = '" + ADMIN_ID + "'",
  );
  assert(adminRole === 'ADMINISTRATOR', 'the seeded bootstrap user is an administrator');
  checks += 1;

  const adminUsername = queryScalar(
    'SELECT username FROM ' + USER_TABLE + " WHERE id = '" + ADMIN_ID + "'",
  );
  assert(
    adminUsername === env.ADMIN_USERNAME,
    'the seeded administrator uses the configured username',
  );
  checks += 1;

  const hashedPassword = queryScalar(
    'SELECT LEFT(password_hash, 3) FROM ' + USER_TABLE + " WHERE id = '" + ADMIN_ID + "'",
  );
  assert(hashedPassword === '$2a', 'the seeded administrator password is a bcrypt hash');
  checks += 1;

  const distinctPasswordHashes = queryScalar(
    'SELECT COUNT(DISTINCT password_hash) FROM ' + USER_TABLE +
      " WHERE id IN ('" + ADMIN_ID + "', '22222222-2222-4222-8222-222222222222', '33333333-3333-4333-8333-333333333333')",
  );
  assert(distinctPasswordHashes === '3', 'each seeded user has a distinct password hash');
  checks += 1;

  const technicianCount = queryScalar('SELECT COUNT(*) FROM technician');
  assert(technicianCount === '2', 'the seed inserted the two bootstrap technicians');
  checks += 1;

  runSql(
    "INSERT INTO customer (id, full_name, document_number, phone, email) VALUES ('"
      + CUSTOMER_ID + "', 'Ana Gomez', 'DOC-0001', '3000000000', 'ana@example.com');",
  );
  runSql(
    "INSERT INTO vehicle (id, customer_id, plate, vin, brand, model, model_year) VALUES ('"
      + VEHICLE_ID + "', '" + CUSTOMER_ID + "', 'ABC123', 'VIN00000000000001', 'Mazda', '3', 2019);",
  );
  checks += 1;

  assert(
    expectSqlError(
      "INSERT INTO vehicle (id, customer_id, plate, vin, brand, model, model_year) VALUES ('e0000000-0000-4000-8000-000000000002', '"
        + CUSTOMER_ID + "', 'ABC123', 'VIN00000000000002', 'Kia', 'Rio', 2020);",
    ),
    'a duplicate plate is rejected',
  );
  checks += 1;

  assert(
    expectSqlError(
      "INSERT INTO vehicle (id, customer_id, plate, vin, brand, model, model_year) VALUES ('e0000000-0000-4000-8000-000000000003', '"
        + CUSTOMER_ID + "', 'XYZ789', 'VIN00000000000001', 'Kia', 'Rio', 2020);",
    ),
    'a duplicate VIN is rejected',
  );
  checks += 1;

  assert(
    expectSqlError(
      "INSERT INTO vehicle (id, customer_id, plate, vin, brand, model, model_year) VALUES ('e0000000-0000-4000-8000-000000000004', 'c0000000-0000-4000-8000-000000000009', 'QQQ111', 'VIN00000000000004', 'Kia', 'Rio', 2020);",
    ),
    'a vehicle without an existing owner is rejected',
  );
  checks += 1;

  runSql(
    "INSERT INTO service_order (id, order_number, vehicle_id, reported_failure, status, received_at) VALUES ('"
      + ORDER_ONE_ID + "', 'OS-0001', '" + VEHICLE_ID + "', 'Ruido en el motor', 'RECEIVED', NOW());",
  );
  runSql(
    "INSERT INTO service_order (id, order_number, vehicle_id, reported_failure, status, received_at) VALUES ('"
      + ORDER_TWO_ID + "', 'OS-0002', '" + VEHICLE_ID + "', 'Frenos gastados', 'RECEIVED', NOW());",
  );
  runSql(
    "INSERT INTO service_order (id, order_number, vehicle_id, reported_failure, status, received_at) VALUES ('"
      + ORDER_THREE_ID + "', 'OS-0003', '" + VEHICLE_ID + "', 'Revision general', 'RECEIVED', NOW());",
  );
  checks += 1;

  assert(
    expectSqlError(
      "INSERT INTO service_order (id, order_number, vehicle_id, reported_failure, status, received_at) VALUES ('50000000-0000-4000-8000-000000000009', 'OS-0009', '"
        + VEHICLE_ID + "', 'Estado invalido', 'CANCELLED', NOW());",
    ),
    'a service order status outside the lifecycle is rejected',
  );
  checks += 1;

  runSql(
    "INSERT INTO assignment (id, service_order_id, technician_id, is_active, active_marker, assigned_at) VALUES ('60000000-0000-4000-8000-000000000001', '"
      + ORDER_ONE_ID + "', '" + TECHNICIAN_ONE_ID + "', 1, '" + TECHNICIAN_ONE_ID + "', NOW());",
  );
  checks += 1;

  assert(
    expectSqlError(
      "INSERT INTO assignment (id, service_order_id, technician_id, is_active, active_marker, assigned_at) VALUES ('60000000-0000-4000-8000-000000000002', '"
        + ORDER_TWO_ID + "', '" + TECHNICIAN_ONE_ID + "', 1, '" + TECHNICIAN_ONE_ID + "', NOW());",
    ),
    'a second active assignment for the same technician is rejected',
  );
  checks += 1;

  runSql(
    "INSERT INTO assignment (id, service_order_id, technician_id, is_active, active_marker, assigned_at) VALUES ('60000000-0000-4000-8000-000000000003', '"
      + ORDER_TWO_ID + "', '" + TECHNICIAN_TWO_ID + "', 1, '" + TECHNICIAN_TWO_ID + "', NOW());",
  );
  checks += 1;

  assert(
    expectSqlError(
      "INSERT INTO assignment (id, service_order_id, technician_id, is_active, active_marker, assigned_at) VALUES ('60000000-0000-4000-8000-000000000004', '"
        + ORDER_THREE_ID + "', '" + TECHNICIAN_TWO_ID + "', 1, NULL, NOW());",
    ),
    'an active assignment without its active marker is rejected',
  );
  checks += 1;

  runSql(
    "INSERT INTO diagnostic (id, service_order_id, technician_id, finding, component_to_repair) VALUES ('90000000-0000-4000-8000-000000000001', '"
      + ORDER_ONE_ID + "', '" + TECHNICIAN_ONE_ID + "', 'Bujias desgastadas', 'Bujias y cables');",
  );
  runSql(
    "INSERT INTO intervention (id, service_order_id, technician_id, description, labor_hour_count, performed_at) VALUES ('"
      + INTERVENTION_ID + "', '" + ORDER_ONE_ID + "', '" + TECHNICIAN_ONE_ID + "', 'Cambio de bujias', 1.50, NOW());",
  );
  checks += 1;

  assert(
    expectSqlError(
      "INSERT INTO intervention (id, service_order_id, technician_id, description, labor_hour_count, performed_at) VALUES ('70000000-0000-4000-8000-000000000009', '"
        + ORDER_ONE_ID + "', '" + TECHNICIAN_ONE_ID + "', 'Sin horas', 0, NOW());",
    ),
    'an intervention without positive labor hours is rejected',
  );
  checks += 1;

  runSql(
    "INSERT INTO part_usage (id, intervention_id, part_name, quantity) VALUES ('a1000000-0000-4000-8000-000000000001', '"
      + INTERVENTION_ID + "', 'Bujia NGK', 4);",
  );
  checks += 1;

  assert(
    expectSqlError(
      "INSERT INTO part_usage (id, intervention_id, part_name, quantity) VALUES ('a1000000-0000-4000-8000-000000000009', '"
        + INTERVENTION_ID + "', 'Bujia NGK', 0);",
    ),
    'a part usage without positive quantity is rejected',
  );
  checks += 1;

  runSql(
    "INSERT INTO warranty (id, intervention_id, warranty_kind, coverage_month_count, issued_at, expiration_date) VALUES ('"
      + WARRANTY_ID + "', '" + INTERVENTION_ID + "', 'LABOR', 12, '2026-01-01 00:00:00', '2027-01-01 00:00:00');",
  );
  checks += 1;

  assert(
    expectSqlError(
      "INSERT INTO warranty (id, intervention_id, warranty_kind, coverage_month_count, issued_at, expiration_date) VALUES ('80000000-0000-4000-8000-000000000009', '"
        + INTERVENTION_ID + "', 'OTHER', 6, NOW(), NOW());",
    ),
    'a warranty of an unknown kind is rejected',
  );
  checks += 1;

  const validBefore = queryScalar(
    "SELECT COUNT(*) FROM warranty WHERE id = '" + WARRANTY_ID + "' AND expiration_date > '2026-06-01 00:00:00'",
  );
  assert(validBefore === '1', 'the warranty reads as valid before its expiration date');
  checks += 1;

  const validAfter = queryScalar(
    "SELECT COUNT(*) FROM warranty WHERE id = '" + WARRANTY_ID + "' AND expiration_date > '2027-06-01 00:00:00'",
  );
  assert(validAfter === '0', 'the warranty reads as expired after its expiration date');
  checks += 1;

  runSql(
    "INSERT INTO status_transition (id, service_order_id, from_status, to_status, changed_by_user_id) VALUES ('b1000000-0000-4000-8000-000000000001', '"
      + ORDER_ONE_ID + "', 'RECEIVED', 'IN_DIAGNOSIS', '" + ADMIN_ID + "');",
  );
  checks += 1;

  assert(
    expectSqlError(
      "INSERT INTO status_transition (id, service_order_id, from_status, to_status, changed_by_user_id) VALUES ('b1000000-0000-4000-8000-000000000009', '"
        + ORDER_ONE_ID + "', 'RECEIVED', 'IN_DIAGNOSIS', '11111111-1111-4111-8111-111111111119');",
    ),
    'a status transition without an existing author is rejected',
  );
  checks += 1;

  const timelineRows = queryScalar(
    "SELECT COUNT(*) FROM service_order so JOIN intervention i ON i.service_order_id = so.id WHERE so.vehicle_id = '"
      + VEHICLE_ID + "'",
  );
  assert(timelineRows === '1', 'the clinical timeline join returns the vehicle intervention');
  checks += 1;

  const statusCount = queryScalar(
    "SELECT COUNT(*) FROM service_order WHERE status = 'RECEIVED'",
  );
  assert(statusCount === '3', 'the dashboard count per status reads the three open orders');
  checks += 1;

  return checks;
}
