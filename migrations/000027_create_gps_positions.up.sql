CREATE TABLE gps_positions (
    id UUID PRIMARY KEY,
    ship_name     TEXT        NOT NULL,
    latitude      DOUBLE PRECISION NOT NULL,
    longitude     DOUBLE PRECISION NOT NULL,
    speed         DOUBLE PRECISION,
    course        DOUBLE PRECISION,
    received_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_gps_positions_ship_time
ON gps_positions (ship_name, received_at DESC);