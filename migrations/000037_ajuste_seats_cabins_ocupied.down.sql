ALTER TABLE reservations_trips 
  RENAME COLUMN occupied_seats TO occupied_units,
  DROP COLUMN IF EXISTS occupied_cabins;