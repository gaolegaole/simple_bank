-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;


-- for update 解决读脏数据，for no key update 解决读脏数据，并且解决由于外键导致的死锁

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id DESC
LIMIT $1
OFFSET $2;

-- name: CreateAccount :one
INSERT INTO accounts (
  owner, balance , currency
) VALUES (
  $1, $2 ,$3
)
RETURNING *;

-- name: UpdateAccount :one
UPDATE accounts 
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: AddAccountBalance :one
UPDATE accounts 
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;