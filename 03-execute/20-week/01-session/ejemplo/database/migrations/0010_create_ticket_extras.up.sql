-- Optional participations of a ticket. A ticketed passenger with no boarding
-- row is a no-show.
CREATE TABLE IF NOT EXISTS payment (
    reference varchar(32) PRIMARY KEY,
    ticket_number varchar(24) NOT NULL REFERENCES ticket(number),
    date timestamptz NOT NULL DEFAULT now(),
    amount numeric(12,2) NOT NULL CHECK (amount >= 0)
);

CREATE TABLE IF NOT EXISTS baggage (
    tag varchar(32) PRIMARY KEY,
    ticket_number varchar(24) NOT NULL REFERENCES ticket(number),
    weight numeric(6,2) NOT NULL CHECK (weight >= 0),
    status varchar(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS boarding (
    ticket_number varchar(24) PRIMARY KEY REFERENCES ticket(number),
    entry_time timestamptz NOT NULL DEFAULT now(),
    gate varchar(12) NOT NULL,
    presentation_condition varchar(20) NOT NULL
);

CREATE INDEX IF NOT EXISTS ix_payment_ticket ON payment(ticket_number);
CREATE INDEX IF NOT EXISTS ix_baggage_ticket ON baggage(ticket_number);
