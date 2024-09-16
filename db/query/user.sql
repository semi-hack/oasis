-- name: CreateUser :one
INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE (id::bigserial IS NULL OR id = $1) AND (email::varchar IS NULL OR email = $2) LIMIT 1;

-- --name: ListUsers :many
-- SELECT * FROM users ORDER BY id LIMIT $1 OFFSET $2;