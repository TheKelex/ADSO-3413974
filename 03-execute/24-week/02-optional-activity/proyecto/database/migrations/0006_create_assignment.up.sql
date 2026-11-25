-- Create the assignment table: which technician holds which service order.
-- The active_marker column holds the technician id while the assignment is
-- active and is null once it is released. Because active_marker is unique, a
-- second active assignment for the same technician is impossible at the
-- storage level, even under a concurrent retry. The check constraint keeps
-- active_marker consistent with is_active.
CREATE TABLE IF NOT EXISTS assignment (
  id CHAR(36) NOT NULL,
  service_order_id CHAR(36) NOT NULL,
  technician_id CHAR(36) NOT NULL,
  is_active TINYINT(1) NOT NULL,
  active_marker CHAR(36) NULL,
  assigned_at DATETIME NOT NULL,
  released_at DATETIME NULL,
  CONSTRAINT pk_assignment PRIMARY KEY (id),
  CONSTRAINT uq_assignment_active_marker UNIQUE (active_marker),
  CONSTRAINT ck_assignment_active_marker CHECK (
    (is_active = 1 AND active_marker IS NOT NULL AND active_marker = technician_id)
    OR (is_active = 0 AND active_marker IS NULL)
  ),
  CONSTRAINT fk_assignment_service_order FOREIGN KEY (service_order_id)
    REFERENCES service_order (id) ON DELETE CASCADE,
  CONSTRAINT fk_assignment_technician FOREIGN KEY (technician_id)
    REFERENCES technician (id) ON DELETE RESTRICT,
  INDEX ix_assignment_service_order_id (service_order_id),
  INDEX ix_assignment_technician_id (technician_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
