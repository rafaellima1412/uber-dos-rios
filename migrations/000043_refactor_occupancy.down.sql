ALTER TABLE public.reservations_trips DROP CONSTRAINT IF EXISTS exclude_no_overlapping_seats;
ALTER TABLE public.reservations_trips DROP CONSTRAINT IF EXISTS exclude_no_overlapping_cabins;


ALTER TABLE public.reservations_trips 
    DROP COLUMN IF EXISTS passenger_id,
    DROP COLUMN IF EXISTS ticket_id,
    DROP COLUMN IF EXISTS seat_code,
    DROP COLUMN IF EXISTS ship_cabin_id,
    DROP COLUMN IF EXISTS origin_stop_order,
    DROP COLUMN IF EXISTS destination_stop_order,
    DROP COLUMN IF EXISTS reserved_until,
    DROP COLUMN IF EXISTS status;