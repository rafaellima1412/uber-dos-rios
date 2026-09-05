DO $$
BEGIN
    -- Verifica se o tipo 'organization_type' ainda não existe no catálogo do sistema.
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'organization_type') THEN
        -- Se não existir, cria o tipo ENUM.
        CREATE TYPE organization_type AS ENUM (
            'COMPANY',
            'RESELLER'
        );
    END IF;
END$$;