CREATE TABLE cities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    state CHAR(2) NOT NULL,
    ibge_code INTEGER UNIQUE, 
    slug VARCHAR(100) 
);
