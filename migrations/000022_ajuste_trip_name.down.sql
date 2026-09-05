ALTER TABLE reservations
DROP CONSTRAINT IF EXISTS reservations_trip_configurations_id_fkey;

ALTER TABLE reservations
RENAME COLUMN trip_configurations_id TO trip_id;


ALTER TABLE trip_configurations RENAME TO trips;






