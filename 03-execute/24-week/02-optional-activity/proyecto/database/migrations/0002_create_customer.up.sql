-- Create the customer table: the owner of the vehicles the workshop services.
-- The national document number identifies the customer for the workshop and is
-- unique; contact data is personal data and is never written to logs.
CREATE TABLE IF NOT EXISTS customer (
  id CHAR(36) NOT NULL,
  full_name VARCHAR(160) NOT NULL,
  document_number VARCHAR(40) NOT NULL,
  phone VARCHAR(40) NOT NULL,
  email VARCHAR(160) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT pk_customer PRIMARY KEY (id),
  CONSTRAINT uq_customer_document_number UNIQUE (document_number)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
