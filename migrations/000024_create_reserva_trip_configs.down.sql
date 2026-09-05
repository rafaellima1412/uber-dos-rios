DROP TABLE IF EXISTS reservations_trip_instances;

ALTER TABLE reservations
ADD COLUMN trip_configurations_id UUID;

ALTER TABLE reservations
ADD CONSTRAINT reservations_trip_id_fkey
FOREIGN KEY (trip_configurations_id)
REFERENCES trip_configurations(id)
ON DELETE CASCADE;


