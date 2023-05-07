// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: payment_game.sql

package db

import (
	"context"
)

const createPaymentGame = `-- name: CreatePaymentGame :one
INSERT INTO payment_games(
    payment_id,
    game_id
) VALUES (
    $1,  $2
)
RETURNING payment_id, game_id
`

type CreatePaymentGameParams struct {
	PaymentID string `json:"payment_id"`
	GameID    int32  `json:"game_id"`
}

func (q *Queries) CreatePaymentGame(ctx context.Context, arg CreatePaymentGameParams) (PaymentGame, error) {
	row := q.db.QueryRowContext(ctx, createPaymentGame, arg.PaymentID, arg.GameID)
	var i PaymentGame
	err := row.Scan(&i.PaymentID, &i.GameID)
	return i, err
}

const deletePaymentGame = `-- name: DeletePaymentGame :one
DELETE FROM payment_games
WHERE payment_id = $1 AND game_id = $2
    RETURNING payment_id, game_id
`

type DeletePaymentGameParams struct {
	PaymentID string `json:"payment_id"`
	GameID    int32  `json:"game_id"`
}

func (q *Queries) DeletePaymentGame(ctx context.Context, arg DeletePaymentGameParams) (PaymentGame, error) {
	row := q.db.QueryRowContext(ctx, deletePaymentGame, arg.PaymentID, arg.GameID)
	var i PaymentGame
	err := row.Scan(&i.PaymentID, &i.GameID)
	return i, err
}

const getPaymentGames = `-- name: GetPaymentGames :many
SELECT games.name FROM payment_games
                           JOIN games ON payment_games.game_id = games.id
WHERE payment_games.payment_id = $1
`

func (q *Queries) GetPaymentGames(ctx context.Context, paymentID string) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getPaymentGames, paymentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
