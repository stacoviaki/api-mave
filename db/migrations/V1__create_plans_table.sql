CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE plans (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,                 -- Ex: Gratuito, Iniciante, Profissional
    price NUMERIC(10,2) DEFAULT 0.00,
    characteristic JSONB,                       -- Ex: { "projetos": 10, "espaco_disponivel": 5 }
    plan_status BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);