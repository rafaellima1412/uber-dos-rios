ALTER TABLE public.reservations_trips DROP COLUMN IF EXISTS occupied_seats;
ALTER TABLE public.reservations_trips DROP COLUMN IF EXISTS occupied_cabins;
DROP INDEX IF EXISTS public.idx_reservations_trips_occupied_seats;