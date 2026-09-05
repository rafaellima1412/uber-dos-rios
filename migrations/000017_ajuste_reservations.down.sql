ALTER TABLE reservations
DROP COLUMN updated_at;

ALTER TABLE reservations
ADD COLUMN schedule_id UUID NOT NULL REFERENCES schedules(id) ON DELETE CASCADE,
ADD COLUMN seat_quantity INTEGER NOT NULL;