CREATE TABLE IF NOT EXISTS docs (
    id SERIAL PRIMARY KEY,
    data JSONB
);
INSERT INTO docs (data)
SELECT jsonb_build_object('tags', ARRAY['tag' || (seq %100), 'common'])
FROM generate_series(1, 100000) as seq;