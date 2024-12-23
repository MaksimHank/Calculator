CREATE TABLE IF NOT EXISTS results (
    id SERIAL PRIMARY KEY,
    first_operand NUMERIC NOT NULL ,
    operator VARCHAR(1) NOT NULL,
    second_operand NUMERIC NOT NULL,
    result NUMERIC NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

DROP TABLE IF EXISTS results;