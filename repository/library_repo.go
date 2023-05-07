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

type ILibraryRepo interface {
	CreateLibrary(ctx context.Context, input entity.CreateLibraryParams) (error, entity.Library)
	DeleteLibrary(ctx context.Context, input entity.DeleteLibraryParams) (error, entity.Library)
	GetUserLibrary(ctx context.Context, userID string) (error, []entity.Game)
	CheckGameInLibrary(ctx context.Context, input entity.CheckGameInLibraryParams) (error, int64)
}

type LibraryRepo struct {
	sql *entity.Queries
}

func NewLibraryRepo(sql *entity.Queries) ILibraryRepo {
	return &LibraryRepo{
		sql: sql,
	}
}

func (l *LibraryRepo) CreateLibrary(ctx context.Context, input entity.CreateLibraryParams) (error, entity.Library) {
	Library, err := l.sql.CreateLibrary(ctx, input)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return errors.New("Library already existed"), entity.Library{}
			}
		}
		log.Println("err: ", err)
		return errors.New("Failed to add Library"), entity.Library{}
	}
	return nil, Library
}

func (l *LibraryRepo) DeleteLibrary(ctx context.Context, input entity.DeleteLibraryParams) (error, entity.Library) {
	library, err := l.sql.DeleteLibrary(ctx, input)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("Library not found"), entity.Library{}
		}
		return err, entity.Library{}
	}
	return nil, library
}

func (l *LibraryRepo) GetUserLibrary(ctx context.Context, userID string) (error, []entity.Game) {
	library, err := l.sql.GetUserLibrary(ctx, userID)
	if err != nil {
		fmt.Printf(err.Error())
		return errors.New("Cannot get all games"), []entity.Game{}
	}
	return nil, library
}

func (l *LibraryRepo) CheckGameInLibrary(ctx context.Context, input entity.CheckGameInLibraryParams) (error, int64) {
	count, err := l.sql.CheckGameInLibrary(ctx, input)
	if err != nil {
		fmt.Printf(err.Error())
		return errors.New("Cannot get all games"), count
	}
	return nil, count
}
