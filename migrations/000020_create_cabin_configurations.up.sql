CREATE TABLE IF NOT EXISTS cabins (
  id uuid PRIMARY KEY,
  name varchar(100) NOT NULL,           
  bed_type varchar(50) NULL,           
  description text NULL,                
  capacity int NOT NULL,
  created_at timestamp DEFAULT now() NOT NULL,
  updated_at timestamp NULL
);

CREATE TABLE IF NOT EXISTS ship_cabins (
  id uuid PRIMARY KEY,
  ship_id uuid NOT NULL REFERENCES ships(id) ON DELETE CASCADE,
  cabin_id uuid NOT NULL REFERENCES cabins(id) ON DELETE CASCADE,
  created_at timestamp DEFAULT now() NOT NULL,
  updated_at timestamp NULL,
  UNIQUE (ship_id, cabin_id)
);

CREATE INDEX IF NOT EXISTS idx_ship_cabins_ship_id ON ship_cabins(ship_id);
CREATE INDEX IF NOT EXISTS idx_ship_cabins_cabin_id ON ship_cabins(cabin_id);