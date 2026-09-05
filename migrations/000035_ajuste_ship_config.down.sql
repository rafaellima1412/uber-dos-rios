-- 1. Revertendo ship_seats para permitir o mesmo navio com diferentes assentos (par composto)
ALTER TABLE ship_seats DROP CONSTRAINT IF EXISTS ship_seats_ship_id_unique;

ALTER TABLE ship_seats ADD CONSTRAINT ship_seats_ship_id_seat_id_key UNIQUE (ship_id, seat_id);