ALTER TABLE IF EXISTS ships
  DROP CONSTRAINT IF EXISTS fk_ships_seat_configuration;

ALTER TABLE IF EXISTS ships
  DROP COLUMN IF EXISTS seat_configuration_id;

ALTER TABLE seat_configurations RENAME TO seats;

ALTER TABLE IF EXISTS seats
  ADD COLUMN IF NOT EXISTS is_active BOOLEAN NOT NULL DEFAULT true,
  DROP COLUMN IF EXISTS nivel;

CREATE TABLE IF NOT EXISTS ship_seats (
  id uuid PRIMARY KEY,
  ship_id uuid NOT NULL REFERENCES ships(id) ON DELETE CASCADE,
  seat_id uuid NOT NULL REFERENCES seats(id) ON DELETE CASCADE,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE,
  UNIQUE (ship_id, seat_id)
);

CREATE INDEX IF NOT EXISTS idx_ship_seats_ship_id ON ship_seats(ship_id);
CREATE INDEX IF NOT EXISTS idx_ship_seats_seat_id ON ship_seats(seat_id)