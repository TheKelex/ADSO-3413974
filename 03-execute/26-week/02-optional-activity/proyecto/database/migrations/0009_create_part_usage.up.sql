-- Create the part usage table: one part consumed inside an intervention. The
-- quantity is strictly positive, and the rows disappear with their intervention
-- because a part usage has no meaning on its own.
CREATE TABLE IF NOT EXISTS part_usage (
  id CHAR(36) NOT NULL,
  intervention_id CHAR(36) NOT NULL,
  part_name VARCHAR(160) NOT NULL,
  quantity INT NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT pk_part_usage PRIMARY KEY (id),
  CONSTRAINT ck_part_usage_quantity CHECK (quantity > 0),
  CONSTRAINT fk_part_usage_intervention FOREIGN KEY (intervention_id)
    REFERENCES intervention (id) ON DELETE CASCADE,
  INDEX ix_part_usage_intervention_id (intervention_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
