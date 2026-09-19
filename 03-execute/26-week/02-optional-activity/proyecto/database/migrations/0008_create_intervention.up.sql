-- Create the intervention table: one physical action executed on the vehicle
-- inside a service order, with the labor hours it consumed. The labor hour
-- count is strictly positive.
CREATE TABLE IF NOT EXISTS intervention (
  id CHAR(36) NOT NULL,
  service_order_id CHAR(36) NOT NULL,
  technician_id CHAR(36) NOT NULL,
  description TEXT NOT NULL,
  labor_hour_count DECIMAL(6,2) NOT NULL,
  performed_at DATETIME NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT pk_intervention PRIMARY KEY (id),
  CONSTRAINT ck_intervention_labor_hour_count CHECK (labor_hour_count > 0),
  CONSTRAINT fk_intervention_service_order FOREIGN KEY (service_order_id)
    REFERENCES service_order (id) ON DELETE CASCADE,
  CONSTRAINT fk_intervention_technician FOREIGN KEY (technician_id)
    REFERENCES technician (id) ON DELETE RESTRICT,
  INDEX ix_intervention_service_order_id (service_order_id),
  INDEX ix_intervention_performed_at (performed_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
