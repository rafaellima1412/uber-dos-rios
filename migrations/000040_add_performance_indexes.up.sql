CREATE INDEX IF NOT EXISTS idx_trips_ship_date ON trips (ship_id, departure_at, arrival_at);
CREATE INDEX IF NOT EXISTS idx_trips_route_departure ON trips (route_id, departure_at);
CREATE INDEX IF NOT EXISTS idx_terminals_city ON terminals (city_id);
CREATE INDEX IF NOT EXISTS idx_ships_org ON ships (organization_id);
CREATE INDEX IF NOT EXISTS idx_trip_configurations_range ON trip_configurations USING GIST (daterange(start_date, expiration_date));
CREATE INDEX IF NOT EXISTS idx_trip_configs_ship_lookup ON public.trip_configurations (ship_id, route_id);
CREATE INDEX IF NOT EXISTS idx_schedules_sequence ON public.schedules (route_id, stop_order, terminal_id);