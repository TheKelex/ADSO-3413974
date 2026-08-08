CREATE TABLE IF NOT EXISTS aircraft (
    registration varchar(16) PRIMARY KEY,
    model varchar(80) NOT NULL,
    capacity integer NOT NULL CHECK (capacity > 0)
);
