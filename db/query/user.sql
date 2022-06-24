-- name: CreateUser :one
INSERT INTO users (
  name,
  email,
  photo,
  verified,
  password,
  role,
  updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
set name = $2,
email = $3,
photo = $4,
verified = $5,
password = $6,
role = $7,
updated_at = $8
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;