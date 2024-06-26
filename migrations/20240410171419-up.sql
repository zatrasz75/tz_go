
-- +migrate Up
CREATE TABLE IF NOT EXISTS people (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255)
    );

CREATE TABLE IF NOT EXISTS cars (
    id SERIAL PRIMARY KEY,
    regNum VARCHAR(255) NOT NULL,
    mark VARCHAR(255) NOT NULL,
    model VARCHAR(255) NOT NULL,
    year INTEGER,
    owner_id INTEGER,
    FOREIGN KEY (owner_id) REFERENCES people(id)
    );

-- +migrate Down
DROP TABLE IF EXISTS cars;
DROP TABLE IF EXISTS people;
