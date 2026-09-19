-- Bootstrap seed: the first administrator plus two technician profiles, so the
-- workshop is operable from the first deployment.
--
-- This file contains no credential. Every value arrives as a MySQL session
-- variable that the caller sets in the same session before piping this file:
--   @admin_username, @admin_full_name, @admin_password_hash,
--   @technician_one_username, @technician_one_full_name,
--   @technician_one_specialty, @technician_one_password_hash,
--   @technician_two_username, @technician_two_full_name,
--   @technician_two_specialty, @technician_two_password_hash
-- The validation runner reads them from .env.example; a real deployment reads
-- them from its own environment. @admin_password_hash is a bcrypt hash: the
-- plain text password never reaches the database.
--
-- The identifiers below are fixed so the seed is idempotent and assertions can
-- address the bootstrap rows deterministically.

INSERT INTO `user` (id, username, password_hash, role, full_name, created_at)
VALUES (
  '11111111-1111-4111-8111-111111111111',
  @admin_username,
  @admin_password_hash,
  'ADMINISTRATOR',
  @admin_full_name,
  CURRENT_TIMESTAMP
)
ON DUPLICATE KEY UPDATE
  password_hash = @admin_password_hash,
  full_name = @admin_full_name;

INSERT INTO `user` (id, username, password_hash, role, full_name, created_at)
VALUES (
  '22222222-2222-4222-8222-222222222222',
  @technician_one_username,
  @technician_one_password_hash,
  'TECHNICIAN',
  @technician_one_full_name,
  CURRENT_TIMESTAMP
)
ON DUPLICATE KEY UPDATE
  password_hash = @technician_one_password_hash,
  full_name = @technician_one_full_name;

INSERT INTO `user` (id, username, password_hash, role, full_name, created_at)
VALUES (
  '33333333-3333-4333-8333-333333333333',
  @technician_two_username,
  @technician_two_password_hash,
  'TECHNICIAN',
  @technician_two_full_name,
  CURRENT_TIMESTAMP
)
ON DUPLICATE KEY UPDATE
  password_hash = @technician_two_password_hash,
  full_name = @technician_two_full_name;

INSERT INTO technician (id, user_id, specialty, created_at)
VALUES (
  'aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaa1',
  '22222222-2222-4222-8222-222222222222',
  @technician_one_specialty,
  CURRENT_TIMESTAMP
)
ON DUPLICATE KEY UPDATE specialty = @technician_one_specialty;

INSERT INTO technician (id, user_id, specialty, created_at)
VALUES (
  'aaaaaaaa-aaaa-4aaa-8aaa-aaaaaaaaaaa2',
  '33333333-3333-4333-8333-333333333333',
  @technician_two_specialty,
  CURRENT_TIMESTAMP
)
ON DUPLICATE KEY UPDATE specialty = @technician_two_specialty;
