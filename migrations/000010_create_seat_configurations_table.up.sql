CREATE TABLE seat_configurations (
    id UUID PRIMARY KEY,
    ship_id UUID NOT NULL REFERENCES ships(id) ON DELETE CASCADE,
    class VARCHAR(20) NOT NULL,
    label VARCHAR(10) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'AVAILABLE',
    position VARCHAR(10) NOT NULL,
    reserved_until TIMESTAMP NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NULL,
    UNIQUE(ship_id, label)
);