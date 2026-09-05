ALTER TABLE reservations_trips RENAME COLUMN occupied_units TO occupied_seats;
ALTER TABLE reservations_trips ADD COLUMN occupied_cabins JSONB;
