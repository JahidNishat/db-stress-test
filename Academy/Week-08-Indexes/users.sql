DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT,
    age INT
);

INSERT INTO users (name, age)
SELECT 'User' || generate_series(1,1000000), (random() * 100)::int;