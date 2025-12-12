-- name: CreateAccount :one
INSERT INTO "accounts" (
  name, slug, domain, domain_verified
)VALUES(
  $1, $2, $3, $4
) RETURNING *;


-- name: RegisterUser :one
INSERT INTO "users" (
  email, name, created_by
) VALUES (
  $1, $2, '00000000-0000-0000-0000-000000000000'
) RETURNING *;

-- name: InsertPassword :one
INSERT INTO "passwords" (
  hashed_password, created_by
) VALUES (
  $1, $2
) RETURNING *;

-- name: FindUserByEmail :one
SELECT * FROM "users" WHERE email = $1 AND deleted_at IS NULL;

-- name: GetPasswordForUser :one
SELECT * FROM "passwords" WHERE created_by = $1 AND deleted_at IS NULL;