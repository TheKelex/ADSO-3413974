CREATE TABLE IF NOT EXISTS passenger (
    document varchar(40) PRIMARY KEY,
    name varchar(160) NOT NULL,
    birth_date date NOT NULL
);
