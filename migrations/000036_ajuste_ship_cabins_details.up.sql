-- 1. Adiciona a coluna jsonb com um valor padrão vazio para não quebrar registros existentes
ALTER TABLE public.ship_cabins 
ADD COLUMN IF NOT EXISTS details jsonb NOT NULL DEFAULT '{}'::jsonb;

-- 2. Cria o índice GIN para permitir buscas rápidas dentro das chaves do JSON
-- (O GIN é essencial para consultas como: details ? 'andar')
CREATE INDEX IF NOT EXISTS idx_ship_cabins_details_gin ON public.ship_cabins USING gin (details);
