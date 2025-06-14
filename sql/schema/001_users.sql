-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    api_key TEXT UNIQUE NOT NULL DEFAULT (
        encode(digest(gen_random_uuid()::text, 'sha256'), 'hex')
    )
);


-- +goose Down
DROP TABLE IF EXISTS users;