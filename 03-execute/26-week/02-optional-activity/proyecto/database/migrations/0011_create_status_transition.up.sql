-- Create the status transition table: the audit record of every service order
-- status change, with the status it came from, the status it moved to, the user
-- who made the change and when. The row is written in the same transaction as
-- the status change, so the history can never disagree with the order.
CREATE TABLE IF NOT EXISTS status_transition (
  id CHAR(36) NOT NULL,
  service_order_id CHAR(36) NOT NULL,
  from_status VARCHAR(40) NOT NULL,
  to_status VARCHAR(40) NOT NULL,
  changed_by_user_id CHAR(36) NOT NULL,
  changed_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT pk_status_transition PRIMARY KEY (id),
  CONSTRAINT fk_status_transition_service_order FOREIGN KEY (service_order_id)
    REFERENCES service_order (id) ON DELETE CASCADE,
  CONSTRAINT fk_status_transition_user FOREIGN KEY (changed_by_user_id)
    REFERENCES `user` (id) ON DELETE RESTRICT,
  INDEX ix_status_transition_service_order_id (service_order_id, changed_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
