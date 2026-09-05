CREATE TABLE IF NOT EXISTS public.bagagens (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    bilhete_id uuid NOT NULL REFERENCES public.tickets(id) ON DELETE CASCADE,
    numero_etiqueta varchar(50) UNIQUE NOT NULL, -- O número do lacre/selo
    descricao varchar(100),                      -- Ex: "Mala azul", "Caixa de isopor com peixe"
    peso_kg numeric(10, 2) NOT NULL DEFAULT 0,
    is_hand_baggage bool DEFAULT false, -- Se for mão (até 10kg), se não (despachada até 30kg)
    status varchar(20) DEFAULT 'DESPACHADA', -- DESPACHADA, A_BORDO, ENTREGUE, EXTRAVIADA
    data_despacho timestamptz DEFAULT now(),
    data_entrega timestamptz,
    
    CONSTRAINT check_peso_positivo CHECK (peso_kg >= 0)
);

CREATE INDEX idx_bagagens_bilhete ON public.bagagens(bilhete_id);
CREATE INDEX idx_bagagens_etiqueta ON public.bagagens(numero_etiqueta);