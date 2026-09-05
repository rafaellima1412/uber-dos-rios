ALTER TABLE ships
DROP CONSTRAINT IF EXISTS fk_ships_seat_configuration;

ALTER TABLE ships
DROP COLUMN IF EXISTS seat_configuration_id;

ALTER TABLE ships
DROP CONSTRAINT IF EXISTS fk_ships_organization;

ALTER TABLE ships
DROP COLUMN IF EXISTS organization_id;