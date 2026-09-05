-- DOWN: rename trips back to trip_instances

-- foreign keys
ALTER TABLE public.trips
RENAME CONSTRAINT fk_trips_trip_configurations
TO fk_trip_configurations_template;

ALTER TABLE public.trips
RENAME CONSTRAINT fk_trips_route
TO fk_trip_instance_route;

ALTER TABLE public.trips
RENAME CONSTRAINT fk_trips_ship
TO fk_trip_instance_ship;

-- primary key
ALTER TABLE public.trips
RENAME CONSTRAINT trips_trip_pkey
TO trip_instances_pkey;

-- table
ALTER TABLE public.trips
RENAME TO trip_instances;
