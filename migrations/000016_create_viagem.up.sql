CREATE TABLE trips (
    id UUID PRIMARY KEY,
    ship_id UUID NOT NULL REFERENCES ships(id) ON DELETE RESTRICT,
    route_id UUID NOT NULL REFERENCES routes(id) ON DELETE RESTRICT,
    recurrence VARCHAR(20) NOT NULL, -- e.g., 'daily', 'weekly', 'monthly'
    expiration_date DATE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE,
    departure_date DATE NOT NULL, --saída da viajem
    arrival_date DATE NOT NULL --chegada da viajem
);

CREATE TABLE reservations (
    id UUID PRIMARY KEY,
    trip_id UUID NOT NULL REFERENCES trips(id) ON DELETE CASCADE,
    schedule_id UUID NOT NULL REFERENCES schedules(id) ON DELETE CASCADE, --escala da viagem
    seat_quantity INTEGER NOT NULL,
    occupied_seats JSONB, -- e.g., seat numbers or identifiers ocupied
    reservation_date DATE NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE SET NULL,
    status VARCHAR(20) NOT NULL, -- e.g., 'confirmed', 'canceled'
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE
);


