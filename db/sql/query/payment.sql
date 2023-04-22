-- name: CreatePayment :one
INSERT INTO payments(
    id, user_id, price
)
VALUES (
           $1,  $2,  $3
       )
    RETURNING *;

-- name: GetUserPayment :many
SELECT *
FROM payments
WHERE user_id = $1;

-- name: DeletePayment :one
DELETE FROM libraries
WHERE user_id = $1 AND game_id = $2
    RETURNING *;
