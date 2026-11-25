-- Create the diagnostic table: the technical evaluation of a service order,
-- written by the technician who holds the active assignment of that order.
CREATE TABLE IF NOT EXISTS diagnostic (
  id CHAR(36) NOT NULL,
  service_order_id CHAR(36) NOT NULL,
  technician_id CHAR(36) NOT NULL,
  finding TEXT NOT NULL,
  component_to_repair TEXT NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT pk_diagnostic PRIMARY KEY (id),
  CONSTRAINT fk_diagnostic_service_order FOREIGN KEY (service_order_id)
    REFERENCES service_order (id) ON DELETE CASCADE,
  CONSTRAINT fk_diagnostic_technician FOREIGN KEY (technician_id)
    REFERENCES technician (id) ON DELETE RESTRICT,
  INDEX ix_diagnostic_service_order_id (service_order_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
