-- Composite identity (flight_number + departure_date); airport plays two
-- roles (origin and destination) and they must differ.
CREATE TABLE IF NOT EXISTS flight (
    flight_number varchar(12) NOT NULL,
    departure_date date NOT NULL,
    scheduled_time time NOT NULL,
    origin_airport_code varchar(8) NOT NULL REFERENCES airport(code),
    destination_airport_code varchar(8) NOT NULL REFERENCES airport(code),
    aircraft_registration varchar(16) NOT NULL REFERENCES aircraft(registration),
    PRIMARY KEY (flight_number, departure_date),
    CONSTRAINT chk_flight_distinct_airports CHECK (origin_airport_code <> destination_airport_code)
);
CREATE INDEX IF NOT EXISTS ix_flight_route ON flight(origin_airport_code, destination_airport_code);
