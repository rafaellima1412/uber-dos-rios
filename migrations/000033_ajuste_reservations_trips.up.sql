-- UP: rename reservations_trip_instances to reservations_trips

ALTER TABLE public.reservations_trip_instances
RENAME TO reservations_trips;

-- primary key
ALTER TABLE public.reservations_trips
RENAME CONSTRAINT reservation_trip_instances_pkey
TO reservations_trips_pkey;

-- foreign keys
ALTER TABLE public.reservations_trips
RENAME CONSTRAINT reservation_trip_configurations_reservation_fkey
TO fk_reservations_trips_reservation;

ALTER TABLE public.reservations_trips
RENAME CONSTRAINT reservation_trip_instances_trip_config_fkey
TO fk_reservations_trips_trip;
