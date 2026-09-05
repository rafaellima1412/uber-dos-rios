 ALTER TABLE terminals
  ADD COLUMN city varchar(50),
  DROP COLUMN city_id;