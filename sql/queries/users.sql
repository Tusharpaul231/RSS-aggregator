-- name: CreateUser :one
-- This query creates a new user in the database.
INSERT INTO users (id, username, email, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW())
RETURNING *;