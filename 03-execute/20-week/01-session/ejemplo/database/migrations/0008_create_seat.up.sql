-- seat is weak to aircraft: identified by seat_number + aircraft_registration.
CREATE TABLE IF NOT EXISTS seat (
    seat_number varchar(8) NOT NULL,
    aircraft_registration varchar(16) NOT NULL REFERENCES aircraft(registration),
    "row" integer NOT NULL,
    location varchar(12) NOT NULL,
    PRIMARY KEY (seat_number, aircraft_registration)
);
