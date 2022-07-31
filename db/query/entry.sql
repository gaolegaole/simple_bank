-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;


-- name: ListEntries :many
SELECT * FROM entries
ORDER BY id DESC
LIMIT $1
OFFSET $2;

-- name: CreateEntry :one
INSERT INTO entries (
  account_id , amount
) VALUES (
  $1, $2
)
RETURNING *;

-- name: ListEntriesByAccount :many
SELECT * FROM entries
WHERE account_id = $3
ORDER BY id DESC
LIMIT $1
OFFSET $2;