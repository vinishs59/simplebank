-- name: CreateTransfer :one
INSERT INTO transfers (
  id,
  from_account,
  to_account,
  amount
) VALUES (
  $1, $2 , $3, $4
)
RETURNING *;


-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers
WHERE 
    from_account = $1 OR
    to_account = $2
ORDER BY id
LIMIT $3
OFFSET $4;