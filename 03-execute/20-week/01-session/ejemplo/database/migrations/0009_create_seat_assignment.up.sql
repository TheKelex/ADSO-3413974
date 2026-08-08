-- Associative table for the ternary seat-passenger-flight assignment.
-- A seat may be assigned at most once on a given flight.
CREATE TABLE IF NOT EXISTS seat_assignment (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    flight_number varchar(12) NOT NULL,
    departure_date date NOT NULL,
    seat_number varchar(8) NOT NULL,
    aircraft_registration varchar(16) NOT NULL,
    passenger_document varchar(40) NOT NULL REFERENCES passenger(document),
    FOREIGN KEY (flight_number, departure_date) REFERENCES flight(flight_number, departure_date),
    FOREIGN KEY (seat_number, aircraft_registration) REFERENCES seat(seat_number, aircraft_registration),
    CONSTRAINT uq_seat_assignment_per_flight UNIQUE (flight_number, departure_date, seat_number, aircraft_registration)
);
