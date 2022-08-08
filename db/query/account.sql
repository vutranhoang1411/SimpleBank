-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY owner;

-- name: CreateAccount :one
INSERT INTO accounts (
  owner,balance,currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;

-- name: UpdateAccount :one
UPDATE  accounts
set owner = $2, balance = $3, currency = $4
WHERE id = $1
RETURNING *;