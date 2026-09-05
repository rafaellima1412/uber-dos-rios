-- 1. Ajustando ship_seats para garantir apenas UM layout por navio
ALTER TABLE ship_seats DROP CONSTRAINT IF EXISTS ship_seats_ship_id_seat_id_key;
-- Adicionamos a nova constraint que proíbe o mesmo ship_id de ter duas linhas
ALTER TABLE ship_seats ADD CONSTRAINT ship_seats_ship_id_unique UNIQUE (ship_id);

-- A constraint UNIQUE (ship_id, cabin_id) já existe no CREATE original, 
