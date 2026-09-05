DROP INDEX IF EXISTS public.idx_ship_cabins_details_gin;

ALTER TABLE public.ship_cabins 
DROP COLUMN IF EXISTS details;