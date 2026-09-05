CREATE TABLE ships (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    imo VARCHAR(20) UNIQUE NOT NULL,
    passenger_capacity INT NOT NULL,
    weight_capacity NUMERIC(10,2) NOT NULL,
    total_seats INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NULL
);