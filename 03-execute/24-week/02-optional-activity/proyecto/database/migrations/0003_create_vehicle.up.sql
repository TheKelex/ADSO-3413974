-- Create the vehicle table: the automobile under service, owned by exactly one
-- customer. The plate and the VIN are unique across the workshop, which is what
-- makes a duplicate registration impossible at the storage level.
CREATE TABLE IF NOT EXISTS vehicle (
  id CHAR(36) NOT NULL,
  customer_id CHAR(36) NOT NULL,
  plate VARCHAR(16) NOT NULL,
  vin VARCHAR(32) NOT NULL,
  brand VARCHAR(80) NOT NULL,
  model VARCHAR(80) NOT NULL,
  model_year SMALLINT NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT pk_vehicle PRIMARY KEY (id),
  CONSTRAINT uq_vehicle_plate UNIQUE (plate),
  CONSTRAINT uq_vehicle_vin UNIQUE (vin),
  CONSTRAINT ck_vehicle_model_year CHECK (model_year > 1900),
  CONSTRAINT fk_vehicle_customer FOREIGN KEY (customer_id)
    REFERENCES customer (id) ON DELETE RESTRICT,
  INDEX ix_vehicle_customer_id (customer_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
