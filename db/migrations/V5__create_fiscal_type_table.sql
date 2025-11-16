CREATE TABLE fiscal_type (
    id SERIAL PRIMARY KEY,
    fiscal_name VARCHAR(255) NOT NULL
);

INSERT INTO fiscal_type (fiscal_name) VALUES
('Mercadoria para Revenda'),
('Matéria-prima'),
('Embalagem'),
('Produto em Processo'),
('Produto Acabado'),
('Subproduto'),
('Produto Intermediário'),
('Material de Uso e Consumo'),
('Ativo Imobilizado'),
('Serviços'),
('Outros Insumos'),
('Outros');
