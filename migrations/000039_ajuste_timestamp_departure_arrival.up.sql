ALTER TABLE public.trips 
    ALTER COLUMN departure_at TYPE timestamptz USING departure_at::timestamptz,
    ALTER COLUMN arrival_at TYPE timestamptz USING arrival_at::timestamptz;