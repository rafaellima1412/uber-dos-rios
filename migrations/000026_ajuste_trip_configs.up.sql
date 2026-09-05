ALTER TABLE trip_configurations
 DROP COLUMN arrival_date,
 DROP COLUMN departure_date;

ALTER TABLE trip_configurations
 ADD COLUMN start_date DATE,
 ADD COLUMN departure_time TIME,
 ADD COLUMN duration_days INTEGER;



