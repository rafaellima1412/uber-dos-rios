DROP INDEX IF EXISTS idx_schedules_route_order;

ALTER TABLE schedules 
DROP COLUMN IF EXISTS stop_order;

