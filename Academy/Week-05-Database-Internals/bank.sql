-- DROP TABLE IF EXISTS accounts;

CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    balance INT NOT NULL DEFAULT 0,
    version INT NOT NULL DEFAULT 1
);

INSERT INTO accounts (id, name, balance, version) VALUES (1, 'Alice', 100, 1);
INSERT INTO accounts (id, name, balance, version) VALUES (2, 'Bob', 50, 1);