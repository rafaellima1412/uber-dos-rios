-- DOWN: rename reservations_trips back to reservations_trip_instances

-- foreign keys
ALTER TABLE public.reservations_trips
RENAME CONSTRAINT fk_reservations_trips_reservation
TO reservation_trip_configurations_reservation_fkey;

ALTER TABLE public.reservations_trips
RENAME CONSTRAINT fk_reservations_trips_trip
TO reservation_trip_instances_trip_config_fkey;

-- primary key
ALTER TABLE public.reservations_trips
RENAME CONSTRAINT reservations_trips_pkey
TO reservation_trip_instances_pkey;

-- table
ALTER TABLE public.reservations_trips
RENAME TO reservations_trip_instances;
