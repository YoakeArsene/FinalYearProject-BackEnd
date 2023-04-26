// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: game.sql

package db

import (
	"context"

	"github.com/lib/pq"
)

const createGame = `-- name: CreateGame :one
INSERT INTO games(
    id,
    name,
    surname,
    price,
    "desc",
    link,
    release,
    platforms,
    genre,
    developers,
    publishers,
    "inCart",
    selected,
    "isHovered",
    "isLiked",
    rating,
    cover,
    footage
) VALUES (
    $1,  $2,  $3,  $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17,
    CAST($18 AS VARCHAR[])
)
RETURNING id, name, surname, price, "desc", link, release, platforms, genre, developers, publishers, "inCart", selected, "isHovered", "isLiked", rating, cover, footage
`

type CreateGameParams struct {
	ID         int32    `json:"id"`
	Name       string   `json:"name"`
	Surname    string   `json:"surname"`
	Price      string   `json:"price"`
	Desc       string   `json:"desc"`
	Link       string   `json:"link"`
	Release    string   `json:"release"`
	Platforms  string   `json:"platforms"`
	Genre      string   `json:"genre"`
	Developers string   `json:"developers"`
	Publishers string   `json:"publishers"`
	InCart     bool     `json:"inCart"`
	Selected   bool     `json:"selected"`
	IsHovered  bool     `json:"isHovered"`
	IsLiked    bool     `json:"isLiked"`
	Rating     int32    `json:"rating"`
	Cover      string   `json:"cover"`
	Column18   []string `json:"column_18"`
}

func (q *Queries) CreateGame(ctx context.Context, arg CreateGameParams) (Game, error) {
	row := q.db.QueryRowContext(ctx, createGame,
		arg.ID,
		arg.Name,
		arg.Surname,
		arg.Price,
		arg.Desc,
		arg.Link,
		arg.Release,
		arg.Platforms,
		arg.Genre,
		arg.Developers,
		arg.Publishers,
		arg.InCart,
		arg.Selected,
		arg.IsHovered,
		arg.IsLiked,
		arg.Rating,
		arg.Cover,
		pq.Array(arg.Column18),
	)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Surname,
		&i.Price,
		&i.Desc,
		&i.Link,
		&i.Release,
		&i.Platforms,
		&i.Genre,
		&i.Developers,
		&i.Publishers,
		&i.InCart,
		&i.Selected,
		&i.IsHovered,
		&i.IsLiked,
		&i.Rating,
		&i.Cover,
		pq.Array(&i.Footage),
	)
	return i, err
}

const deleteGame = `-- name: DeleteGame :one
DELETE FROM games
WHERE id = $1
    RETURNING id, name, surname, price, "desc", link, release, platforms, genre, developers, publishers, "inCart", selected, "isHovered", "isLiked", rating, cover, footage
`

func (q *Queries) DeleteGame(ctx context.Context, id int32) (Game, error) {
	row := q.db.QueryRowContext(ctx, deleteGame, id)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Surname,
		&i.Price,
		&i.Desc,
		&i.Link,
		&i.Release,
		&i.Platforms,
		&i.Genre,
		&i.Developers,
		&i.Publishers,
		&i.InCart,
		&i.Selected,
		&i.IsHovered,
		&i.IsLiked,
		&i.Rating,
		&i.Cover,
		pq.Array(&i.Footage),
	)
	return i, err
}

const getAllGames = `-- name: GetAllGames :many
Select id, name, surname, price, "desc", link, release, platforms, genre, developers, publishers, "inCart", selected, "isHovered", "isLiked", rating, cover, footage FROM games
`

func (q *Queries) GetAllGames(ctx context.Context) ([]Game, error) {
	rows, err := q.db.QueryContext(ctx, getAllGames)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Game
	for rows.Next() {
		var i Game
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Surname,
			&i.Price,
			&i.Desc,
			&i.Link,
			&i.Release,
			&i.Platforms,
			&i.Genre,
			&i.Developers,
			&i.Publishers,
			&i.InCart,
			&i.Selected,
			&i.IsHovered,
			&i.IsLiked,
			&i.Rating,
			&i.Cover,
			pq.Array(&i.Footage),
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateGame = `-- name: UpdateGame :one
UPDATE games
SET  name = $1,
     surname = $2,
     price = $3,
     "desc" = $4,
     link = $5,
     release = $6,
     platforms = $7,
     genre = $8,
     developers = $9,
     publishers = $10,
     "inCart" = $11,
     selected = $12,
     "isHovered" = $13,
     "isLiked" = $14,
     rating = $15,
     cover = $16,
     footage = CAST($17 AS VARCHAR[])
WHERE id = $18
    RETURNING id, name, surname, price, "desc", link, release, platforms, genre, developers, publishers, "inCart", selected, "isHovered", "isLiked", rating, cover, footage
`

type UpdateGameParams struct {
	Name       string   `json:"name"`
	Surname    string   `json:"surname"`
	Price      string   `json:"price"`
	Desc       string   `json:"desc"`
	Link       string   `json:"link"`
	Release    string   `json:"release"`
	Platforms  string   `json:"platforms"`
	Genre      string   `json:"genre"`
	Developers string   `json:"developers"`
	Publishers string   `json:"publishers"`
	InCart     bool     `json:"inCart"`
	Selected   bool     `json:"selected"`
	IsHovered  bool     `json:"isHovered"`
	IsLiked    bool     `json:"isLiked"`
	Rating     int32    `json:"rating"`
	Cover      string   `json:"cover"`
	Column17   []string `json:"column_17"`
	ID         int32    `json:"id"`
}

func (q *Queries) UpdateGame(ctx context.Context, arg UpdateGameParams) (Game, error) {
	row := q.db.QueryRowContext(ctx, updateGame,
		arg.Name,
		arg.Surname,
		arg.Price,
		arg.Desc,
		arg.Link,
		arg.Release,
		arg.Platforms,
		arg.Genre,
		arg.Developers,
		arg.Publishers,
		arg.InCart,
		arg.Selected,
		arg.IsHovered,
		arg.IsLiked,
		arg.Rating,
		arg.Cover,
		pq.Array(arg.Column17),
		arg.ID,
	)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Surname,
		&i.Price,
		&i.Desc,
		&i.Link,
		&i.Release,
		&i.Platforms,
		&i.Genre,
		&i.Developers,
		&i.Publishers,
		&i.InCart,
		&i.Selected,
		&i.IsHovered,
		&i.IsLiked,
		&i.Rating,
		&i.Cover,
		pq.Array(&i.Footage),
	)
	return i, err
}
