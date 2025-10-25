CREATE TABLE contacts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    -- Tipo de contato
    contact_type VARCHAR(20) NOT NULL CHECK (contact_type IN ('individual', 'company')),

    -- Dados gerais
    full_name VARCHAR(255),       -- nome completo da pessoa física
    trade_name VARCHAR(255),      -- nome fantasia (empresa)
    company_name VARCHAR(255),    -- razão social (empresa)

    -- Documentos
    cpf VARCHAR(14),
    rg VARCHAR(20),
    cnpj VARCHAR(18),
    foreing_id VARCHAR(50),
    state_registration VARCHAR(20),
    municipal_registration VARCHAR(20),

    -- Endereço
    postal_code VARCHAR(10),
    street VARCHAR(255),
    number VARCHAR(20),
    complement VARCHAR(100),
    district VARCHAR(100),
    city VARCHAR(100),
    state VARCHAR(100),
    country_id INT REFERENCES countries(id),

    -- Contato
    phone VARCHAR(20),
    mobile VARCHAR(20),
    email VARCHAR(255),
    website VARCHAR(255),

    -- Profissional / Comercial
    job_title VARCHAR(100),
    tags VARCHAR(255),

    -- Controle
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



-- Cria a função usada pelos triggers para atualizar automaticamente o updated_at
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER update_contacts_timestamp
BEFORE UPDATE ON contacts
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();