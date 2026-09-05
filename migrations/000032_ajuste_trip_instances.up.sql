-- UP: rename trip_instances to trips

ALTER TABLE public.trip_instances
RENAME TO trips;

-- primary key
ALTER TABLE public.trips
RENAME CONSTRAINT trip_instances_pkey
TO trips_trip_pkey;

-- foreign keys
ALTER TABLE public.trips
RENAME CONSTRAINT fk_trip_configurations_template
TO fk_trips_trip_configurations;

ALTER TABLE public.trips
RENAME CONSTRAINT fk_trip_instance_route
TO fk_trips_route;

ALTER TABLE public.trips
RENAME CONSTRAINT fk_trip_instance_ship
TO fk_trips_ship;
