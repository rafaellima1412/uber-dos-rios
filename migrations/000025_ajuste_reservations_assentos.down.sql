ALTER TABLE reservations_trip_instances
 DROP COLUMN occupied_seats;

ALTER TABLE reservations
 ADD COLUMN occupied_seats JSONB;

