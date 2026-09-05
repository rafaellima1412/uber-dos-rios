ALTER TABLE schedules 
ADD COLUMN stop_order INTEGER NOT NULL DEFAULT 0;

CREATE INDEX idx_schedules_route_order ON schedules(route_id, stop_order);