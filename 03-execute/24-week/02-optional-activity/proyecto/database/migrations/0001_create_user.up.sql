-- Create the user table: the authentication account and the role of every
-- person who signs in. The password is stored as a bcrypt hash produced by the
-- backend; a plain text password is never persisted.
CREATE TABLE IF NOT EXISTS `user` (
  id CHAR(36) NOT NULL,
  username VARCHAR(80) NOT NULL,
  password_hash VARCHAR(255) NOT NULL,
  role VARCHAR(40) NOT NULL,
  full_name VARCHAR(160) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT pk_user PRIMARY KEY (id),
  CONSTRAINT uq_user_username UNIQUE (username),
  CONSTRAINT ck_user_role CHECK (role IN ('ADMINISTRATOR', 'TECHNICIAN'))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
