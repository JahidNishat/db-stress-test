DROP TABLE IF EXISTS counter;

-- The table we will abuse
CREATE TABLE IF NOT EXISTS counter (
    id SERIAL PRIMARY KEY,
    val INT NOT NULL DEFAULT 0,
    version INT NOT NULL DEFAULT 1
);

-- Reset function to clear data before test
TRUNCATE counter;
INSERT INTO counter (id, val, version) VALUES (1, 0, 1);
