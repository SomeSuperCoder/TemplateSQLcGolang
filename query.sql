-- name: FindAllPlayers :many
SELECT * FROM players;

-- name: InsertItem :one
INSERT INTO items (id, name, value)
VALUES (uuid_generate_v4(), $1, $2)
RETURNING *;
