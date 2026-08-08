-- pgcrypto provides gen_random_uuid() and crypt()/gen_salt() for bcrypt.
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- "user" is a reserved word in PostgreSQL, so it is quoted here and in queries.
CREATE TABLE IF NOT EXISTS "user" (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    username varchar(80) NOT NULL UNIQUE,
    password_hash varchar(255) NOT NULL,
    role varchar(40) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now()
);
