-- Create the warranty table: the coverage issued over one intervention. The
-- expiration date is computed by the application as the issue date plus the
-- coverage in months and is persisted, so validity is read, never recomputed
-- from a rule that may have changed. The intervention cannot be removed while
-- a warranty points at it.
CREATE TABLE IF NOT EXISTS warranty (
  id CHAR(36) NOT NULL,
  intervention_id CHAR(36) NOT NULL,
  warranty_kind VARCHAR(20) NOT NULL,
  coverage_month_count INT NOT NULL,
  issued_at DATETIME NOT NULL,
  expiration_date DATETIME NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT pk_warranty PRIMARY KEY (id),
  CONSTRAINT ck_warranty_kind CHECK (warranty_kind IN ('LABOR', 'PART')),
  CONSTRAINT ck_warranty_coverage_month_count CHECK (coverage_month_count > 0),
  CONSTRAINT fk_warranty_intervention FOREIGN KEY (intervention_id)
    REFERENCES intervention (id) ON DELETE RESTRICT,
  INDEX ix_warranty_intervention_id (intervention_id),
  INDEX ix_warranty_expiration_date (expiration_date)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
