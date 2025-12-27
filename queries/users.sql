-- name: InsertUser :one
INSERT INTO users (
  name,
  email,
  password
) VALUES (
  $1,
  $2,
  crypt($3, gen_salt('bf'))
) RETURNING 
  id,
  name,
  email,
  created_at
;

-- name: UpdateProfile :one
