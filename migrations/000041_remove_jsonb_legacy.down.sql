ALTER TABLE public.reservations_trips ADD COLUMN occupied_seats jsonb DEFAULT '[]'::jsonb;
ALTER TABLE public.reservations_trips ADD COLUMN occupied_cabins jsonb DEFAULT '[]'::jsonb;
CREATE INDEX IF NOT EXISTS idx_reservations_trips_occupied_seats ON public.reservations_trips USING gin (occupied_seats);