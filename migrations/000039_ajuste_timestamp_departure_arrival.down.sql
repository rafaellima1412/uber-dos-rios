ALTER TABLE public.trips 
    ALTER COLUMN departure_at TYPE timestamp WITHOUT TIME ZONE USING departure_at::timestamp WITHOUT TIME ZONE,
    ALTER COLUMN arrival_at TYPE timestamp WITHOUT TIME ZONE USING arrival_at::timestamp WITHOUT TIME ZONE;