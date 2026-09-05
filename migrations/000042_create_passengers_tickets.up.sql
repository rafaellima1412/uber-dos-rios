CREATE TABLE public.passengers (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name varchar(100) NOT NULL,
    last_name varchar(100) NOT NULL,
    document_number varchar(20) NOT NULL, 
    document_type varchar(10) NOT NULL,
    birth_date date NOT NULL,
    created_at timestamptz DEFAULT now()
);

CREATE TABLE public.tickets (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    reservation_id uuid NOT NULL REFERENCES public.reservations(id) ON DELETE CASCADE,
    passenger_id uuid NOT NULL REFERENCES public.passengers(id),
    ticket_number varchar(20) UNIQUE,
    status varchar(20) DEFAULT 'ISSUED',
    valor_tarifa numeric(10, 2) NOT NULL DEFAULT 0,   -- Preço base do trecho
    valor_taxas numeric(10, 2) NOT NULL DEFAULT 0,    -- Taxas de terminais/portos
    valor_seguro numeric(10, 2) NOT NULL DEFAULT 0,   -- Seguro obrigatório
    seguradora_nome varchar(100),
    seguradora_cnpj varchar(20),
    seguradora_apolice varchar(50),
    tipo_desconto varchar(30), -- 'IDOSO_GRATUITO', 'IDOSO_50', 'ESTUDANTE', 'CRIANCA_COLO'
    tem_excesso_bagagem_pago bool DEFAULT false,
    created_at timestamptz DEFAULT now()
);

CREATE INDEX idx_bilhetes_reserva ON public.tickets(reservation_id);
CREATE INDEX idx_bilhetes_passageiro ON public.tickets( passenger_id);