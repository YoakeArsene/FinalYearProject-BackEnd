package repository

import (
	entity "GDN-delivery-management/db/sql"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/lib/pq"
)

type IGameRepo interface {
	CreateGame(ctx context.Context, input entity.CreateGameParams) (error, entity.Game)
	DeleteGame(ctx context.Context, id int32) (error, entity.Game)
	GetAllGames(ctx context.Context) (error, []entity.Game)
	UpdateGame(ctx context.Context, input entity.UpdateGameParams) (error, entity.Game)
}

type GameRepo struct {
	sql *entity.Queries
}

func NewGameRepo(sql *entity.Queries) IGameRepo {
	return &GameRepo{
		sql: sql,
	}
}

func (g *GameRepo) CreateGame(ctx context.Context, input entity.CreateGameParams) (error, entity.Game) {
	game, err := g.sql.CreateGame(ctx, input)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return errors.New("Game already existed"), entity.Game{}
			}
		}
		log.Println("err: ", err)
		return errors.New("Failed to add game"), entity.Game{}
	}
	return nil, game
}

func (g *GameRepo) DeleteGame(ctx context.Context, id int32) (error, entity.Game) {
	game, err := g.sql.DeleteGame(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("game not found"), entity.Game{}
		}
		return err, entity.Game{}
	}
	return nil, game
}

func (g *GameRepo) GetAllGames(ctx context.Context) (error, []entity.Game) {
	items, err := g.sql.GetAllGames(ctx)
	if err != nil {
		fmt.Printf(err.Error())
		return errors.New("Cannot get all games"), []entity.Game{}
	}
	return nil, items
}

func (g *GameRepo) UpdateGame(ctx context.Context, input entity.UpdateGameParams) (error, entity.Game) {
	game, err := g.sql.UpdateGame(ctx, input)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("game not found"), entity.Game{}
		}
		return err, entity.Game{}
	}
	return nil, game
}
