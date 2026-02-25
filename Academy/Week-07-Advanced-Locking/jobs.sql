CREATE TABLE IF NOT EXISTS jobs (
    id SERIAL PRIMARY KEY,
    status TEXT DEFAULT 'pending'
);

INSERT INTO jobs (status) VALUES ('pending'), ('pending'), ('pending');