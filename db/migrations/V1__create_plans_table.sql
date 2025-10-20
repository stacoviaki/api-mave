CREATE TABLE planos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,                 -- Ex: Gratuito, Iniciante, Profissional
    preco NUMERIC(10,2) DEFAULT 0.00,
    caracteristicas JSONB,                             -- Ex: { "projetos": 10, "Espaco disponivel": 5 }
    status_plano BOOLEAN DEFAULT TRUE,
    criado_em TIMESTAMP DEFAULT NOW()
);