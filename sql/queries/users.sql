-- name: CreateUser :one
-- This query creates a new user in the database.
INSERT INTO users (id, username, email, created_at, updated_at, api_key)
VALUES (
    $1,
    $2,
    $3,
    NOW(),
    NOW(),
    encode(digest(random()::text, 'sha256'), 'hex')
)
RETURNING *;

-- name: GetUserByAPIKey :one
-- This query retrieves a user by their ID.
SELECT * FROM users 
WHERE api_key = $1;
