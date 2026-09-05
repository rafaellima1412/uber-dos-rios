ALTER TABLE reservations
DROP CONSTRAINT IF EXISTS reservations_trip_id_fkey;

ALTER TABLE reservations
DROP COLUMN IF EXISTS trip_configurations_id;

CREATE TABLE reservations_trip_instances (
    reservation_id UUID NOT NULL,
    trip_instances_id UUID NOT NULL,

    CONSTRAINT reservation_trip_instances_pkey
        PRIMARY KEY (reservation_id, trip_instances_id),

    CONSTRAINT reservation_trip_configurations_reservation_fkey
        FOREIGN KEY (reservation_id)
        REFERENCES reservations(id)
        ON DELETE CASCADE,

    CONSTRAINT reservation_trip_instances_trip_config_fkey
        FOREIGN KEY (trip_instances_id)
        REFERENCES trip_instances(id)
        ON DELETE CASCADE
);
