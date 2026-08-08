CREATE TABLE IF NOT EXISTS reservation (
    code varchar(24) PRIMARY KEY,
    passenger_document varchar(40) NOT NULL REFERENCES passenger(document),
    date timestamptz NOT NULL DEFAULT now(),
    status varchar(20) NOT NULL
);
CREATE INDEX IF NOT EXISTS ix_reservation_passenger ON reservation(passenger_document);
