ALTER TABLE public.reservations_trips ADD COLUMN passenger_id uuid REFERENCES public.passengers(id);
ALTER TABLE public.reservations_trips ADD COLUMN ticket_id uuid REFERENCES public.tickets(id);
ALTER TABLE public.reservations_trips ADD COLUMN seat_code varchar(10);
ALTER TABLE public.reservations_trips ADD COLUMN ship_cabin_id uuid REFERENCES public.ship_cabins(id);
ALTER TABLE public.reservations_trips ADD COLUMN origin_stop_order integer;
ALTER TABLE public.reservations_trips ADD COLUMN destination_stop_order integer;
ALTER TABLE public.reservations_trips ADD COLUMN IF NOT EXISTS reserved_until timestamptz;
ALTER TABLE public.reservations_trips ADD COLUMN IF NOT EXISTS status varchar(20) DEFAULT 'CONFIRMED';
CREATE EXTENSION IF NOT EXISTS btree_gist;

-- Adiciona a trava física (Constraint) em vez de apenas um índice
ALTER TABLE public.reservations_trips 
ADD CONSTRAINT exclude_no_overlapping_seats
EXCLUDE USING gist (
    trip_instances_id WITH =, 
    seat_code WITH =,
    int4range(origin_stop_order, destination_stop_order) WITH &&
) WHERE (seat_code IS NOT NULL AND status != 'CANCELLED');