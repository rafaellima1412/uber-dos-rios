CREATE TABLE IF NOT EXISTS organizations (
    id UUID PRIMARY KEY,
    org_type organization_type NOT NULL,
    name VARCHAR(255) NOT NULL,
    cnpj VARCHAR(14) UNIQUE,
    cpf VARCHAR(11) UNIQUE,
    email VARCHAR(255),
    phone_number VARCHAR(20),
    address JSONB,
    owner_info JSONB,
    logo_url TEXT,
    custom_login_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT check_document CHECK (
        (cpf IS NOT NULL AND cnpj IS NULL) OR (cnpj IS NOT NULL AND cpf IS NULL)
    )
);