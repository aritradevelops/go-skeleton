-- name: RegisterUser :one
INSERT INTO "users" (
  email, name, created_by
) VALUES (
  $1, $2, '00000000-0000-0000-0000-000000000000'
) RETURNING *;
