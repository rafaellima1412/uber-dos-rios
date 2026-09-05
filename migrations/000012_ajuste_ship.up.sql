ALTER TABLE ships
ADD COLUMN IF NOT EXISTS seat_configuration_id UUID;

ALTER TABLE ships
ADD CONSTRAINT fk_ships_seat_configuration
FOREIGN KEY (seat_configuration_id)
REFERENCES seat_configurations(id)
ON DELETE SET NULL
ON UPDATE CASCADE;

ALTER TABLE ships
ADD COLUMN IF NOT EXISTS organization_id UUID;

ALTER TABLE ships
ADD CONSTRAINT fk_ships_organization
FOREIGN KEY (organization_id)
REFERENCES organizations(id)
ON DELETE SET NULL
ON UPDATE CASCADE;