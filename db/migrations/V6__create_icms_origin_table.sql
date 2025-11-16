CREATE TABLE icms_origin (
    id SERIAL PRIMARY KEY,
    icms_origin_code INTEGER NOT NULL,
    icms_origin_description VARCHAR(500) NOT NULL
);

INSERT INTO icms_origin (icms_origin_code, icms_origin_description) VALUES
(0, 'Nacional, exceto as indicadas nos códigos 3, 4, 5 e 8'),
(1, 'Estrangeira – importação direta, exceto a indicada no código 6'),
(2, 'Estrangeira – adquirida no mercado interno, exceto a indicada no código 7'),
(3, 'Nacional – mercadoria ou bem com Conteúdo de Importação superior a 40% e inferior ou igual a 70%'),
(4, 'Nacional – cuja produção foi feita em conformidade com os processos produtivos básicos (Decreto-Lei nº 288/67, Leis 8.248/91, 8.387/91, 10.176/2001 e 11.484/2007)'),
(5, 'Nacional – mercadoria ou bem com Conteúdo de Importação inferior ou igual a 40%'),
(6, 'Estrangeira – importação direta, sem similar nacional, constante na lista da Resolução CAMEX e gás natural'),
(7, 'Estrangeira – adquirida no mercado interno, sem similar nacional, constante na lista da Resolução CAMEX e gás natural'),
(8, 'Nacional – mercadoria ou bem com Conteúdo de Importação superior a 70%');
