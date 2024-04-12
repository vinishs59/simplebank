-- name: CreateAccount :one
INSERT INTO accounts (
  owner_name,
  balance,
  currency
) VALUES (
  $1, $2 ,$3
)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE user_id = $1 LIMIT 1;

-- name: ListAccount :many
SELECT * FROM accounts
ORDER BY user_id
LIMIT $1
OFFSET $2;


-- name: UpdateAccount :exec
UPDATE accounts SET balance = $2
WHERE user_id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts 
WHERE user_id = $1;