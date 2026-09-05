ALTER TABLE reservations_trips
RENAME COLUMN occupied_units TO occupied_seats;

COMMENT ON COLUMN reservations_trips.occupied_seats IS NULL;
