ALTER TABLE trips RENAME TO trip_configurations;
ALTER TABLE reservations
DROP CONSTRAINT reservations_trip_id_fkey;

ALTER TABLE reservations
RENAME COLUMN trip_id TO trip_configurations_id;

ALTER TABLE reservations
ADD CONSTRAINT reservations_trip_configurations_id_fkey
FOREIGN KEY (trip_configurations_id)
REFERENCES trip_configurations(id)
ON DELETE CASCADE;