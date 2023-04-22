-- name: CreatePaymentGame :one
INSERT INTO payment_games(
    payment_id,
    game_id
) VALUES (
    $1,  $2
)
RETURNING *;

-- name: GetPaymentGames :many
SELECT * FROM payment_games WHERE payment_id = $1;

-- name: DeletePaymentGame :one
DELETE FROM payment_games
WHERE payment_id = $1 AND game_id = $2
    RETURNING *;