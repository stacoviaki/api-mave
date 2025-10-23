CREATE TABLE countries (
    id SERIAL PRIMARY KEY,
    iso_code CHAR(2) UNIQUE NOT NULL,
    country_name VARCHAR(100) NOT NULL,
    phone_code VARCHAR(10),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Inserção automática logo após criar a tabela
INSERT INTO countries (iso_code, country_name, phone_code) VALUES
('BR','Brasil','+55'),
('US','Estados Unidos','+1'),
('CA','Canadá','+1'),
('MX','México','+52'),
('PT','Portugal','+351'),
('ES','Espanha','+34'),
('FR','França','+33'),
('DE','Alemanha','+49'),
('IT','Itália','+39'),
('JP','Japão','+81'),
('CN','China','+86'),
('IN','Índia','+91'),
('AU','Austrália','+61'),
('RU','Rússia','+7'),
('ZA','África do Sul','+27'),
('AR','Argentina','+54'),
('CL','Chile','+56'),
('CO','Colômbia','+57');
