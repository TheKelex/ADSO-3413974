CREATE TABLE IF NOT EXISTS ticket (
    number varchar(24) PRIMARY KEY,
    reservation_code varchar(24) NOT NULL REFERENCES reservation(code),
    issue_date timestamptz NOT NULL DEFAULT now(),
    service_class varchar(20) NOT NULL
);
CREATE INDEX IF NOT EXISTS ix_ticket_reservation ON ticket(reservation_code);
