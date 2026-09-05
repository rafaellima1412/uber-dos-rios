ALTER TABLE reservations
DROP COLUMN occupied_seats;

Alter TABLE reservations_trip_instances
ADD COLUMN occupied_seats JSONB;