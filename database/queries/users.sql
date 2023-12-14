
-- name: GetUserByID :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY email;

-- name: CreateUser :one
INSERT INTO users (
  email, password_hash, created_at, updated_at
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
SET email = ?,
    password_hash = ?,
    updated_at = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;

