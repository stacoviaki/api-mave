CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE products (
    -- UUID
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    -- Informações do Produto
    product_name VARCHAR(60) NOT NULL,
    sale_price NUMERIC(12,5),       
    product_price NUMERIC(12,5),    
    product_reference VARCHAR(30),

    -- Fiscal
    fiscal_type INT REFERENCES fiscal_type(id),
    icms_origin INT REFERENCES icms_origin(id),
    ncm INT REFERENCES ncm(id),
    product_fiscal_type INT REFERENCES product_fiscal_type(id),
    cest INT REFERENCES cest(id),
    nbm INT REFERENCES nbm(id),

    -- Logs
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Função para atualizar automaticamente o updated_at
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger para atualizar updated_at antes de UPDATE
CREATE TRIGGER update_products_timestamp
BEFORE UPDATE ON products
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
