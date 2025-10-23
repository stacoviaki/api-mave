CREATE TABLE countries (
    id SERIAL PRIMARY KEY,
    iso_code CHAR(2) UNIQUE NOT NULL,
    iso_code3 CHAR(3) UNIQUE,
    name_pt VARCHAR(100) NOT NULL,
    name_en VARCHAR(100) NOT NULL,
    phone_code VARCHAR(10),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Inserção automática logo após criar a tabela
INSERT INTO countries (iso_code, iso_code3, name_pt, name_en, phone_code) VALUES
('BR','BRA','Brasil','Brazil','+55'),
('US','USA','Estados Unidos','United States','+1'),
('CA','CAN','Canadá','Canada','+1'),
('MX','MEX','México','Mexico','+52'),
('PT','PRT','Portugal','Portugal','+351'),
('ES','ESP','Espanha','Spain','+34'),
('FR','FRA','França','France','+33'),
('DE','DEU','Alemanha','Germany','+49'),
('IT','ITA','Itália','Italy','+39'),
('JP','JPN','Japão','Japan','+81'),
('CN','CHN','China','China','+86'),
('IN','IND','Índia','India','+91'),
('AU','AUS','Austrália','Australia','+61'),
('RU','RUS','Rússia','Russia','+7'),
('ZA','ZAF','África do Sul','South Africa','+27'),
('AR','ARG','Argentina','Argentina','+54'),
('CL','CHL','Chile','Chile','+56'),
('CO','COL','Colômbia','Colombia','+57');