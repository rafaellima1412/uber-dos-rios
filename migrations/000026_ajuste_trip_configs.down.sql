ALTER TABLE trip_configurations
 DROP COLUMN departure_time,
 DROP COLUMN duration_days,
 DROP COLUMN start_date;

ALTER TABLE trip_configurations
 ADD COLUMN departure_date TIMESTAMP NOT NULL,
 ADD COLUMN arrival_date TIMESTAMP NOT NULL;
