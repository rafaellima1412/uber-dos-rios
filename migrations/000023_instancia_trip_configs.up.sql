CREATE TABLE trip_instances (
	id UUID PRIMARY KEY,
	trip_configurations_id UUID NOT NULL,
	route_id UUID NOT NULL,
	ship_id UUID NOT NULL,
	departure_at TIMESTAMP NOT NULL,
	arrival_at   TIMESTAMP NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE,

	CONSTRAINT fk_trip_configurations_template
		FOREIGN KEY (trip_configurations_id)
		REFERENCES trip_configurations(id),

	CONSTRAINT fk_trip_instance_route
		FOREIGN KEY (route_id)
		REFERENCES routes(id),

	CONSTRAINT fk_trip_instance_ship
		FOREIGN KEY (ship_id)
		REFERENCES ships(id)
);



