-- Create the technician table: the mechanic profile bound to exactly one user
-- account. The unique user reference keeps one account mapped to at most one
-- mechanic profile.
CREATE TABLE IF NOT EXISTS technician (
  id CHAR(36) NOT NULL,
  user_id CHAR(36) NOT NULL,
  specialty VARCHAR(120) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT pk_technician PRIMARY KEY (id),
  CONSTRAINT uq_technician_user_id UNIQUE (user_id),
  CONSTRAINT fk_technician_user FOREIGN KEY (user_id)
    REFERENCES `user` (id) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
