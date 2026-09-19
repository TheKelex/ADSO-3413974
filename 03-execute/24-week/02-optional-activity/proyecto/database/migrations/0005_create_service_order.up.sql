-- Create the service order table: the work order opened at check-in for one
-- vehicle. The status column carries the lifecycle and is constrained to the
-- declared values; the application owns the allowed transitions between them.
CREATE TABLE IF NOT EXISTS service_order (
  id CHAR(36) NOT NULL,
  order_number VARCHAR(32) NOT NULL,
  vehicle_id CHAR(36) NOT NULL,
  reported_failure TEXT NOT NULL,
  status VARCHAR(40) NOT NULL,
  received_at DATETIME NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT pk_service_order PRIMARY KEY (id),
  CONSTRAINT uq_service_order_order_number UNIQUE (order_number),
  CONSTRAINT ck_service_order_status CHECK (
    status IN ('RECEIVED', 'IN_DIAGNOSIS', 'IN_REPAIR', 'READY', 'DELIVERED')
  ),
  CONSTRAINT fk_service_order_vehicle FOREIGN KEY (vehicle_id)
    REFERENCES vehicle (id) ON DELETE RESTRICT,
  INDEX ix_service_order_vehicle_id (vehicle_id),
  INDEX ix_service_order_status (status),
  INDEX ix_service_order_received_at (received_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
