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

type IPaymentGameRepo interface {
	CreatePaymentGame(ctx context.Context, input entity.CreatePaymentGameParams) (error, entity.PaymentGame)
	DeletePaymentGame(ctx context.Context, input entity.DeletePaymentGameParams) (error, entity.PaymentGame)
	GetPaymentGames(ctx context.Context, paymentID string) (error, []string)
}

type PaymentGameRepo struct {
	sql *entity.Queries
}

func NewPaymentGameRepo(sql *entity.Queries) IPaymentGameRepo {
	return &PaymentGameRepo{
		sql: sql,
	}
}

func (pg *PaymentGameRepo) CreatePaymentGame(ctx context.Context, input entity.CreatePaymentGameParams) (error, entity.PaymentGame) {
	paymentGame, err := pg.sql.CreatePaymentGame(ctx, input)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return errors.New("Payment already existed"), entity.PaymentGame{}
			}
		}
		log.Println("err: ", err)
		return errors.New("Failed to add payment"), entity.PaymentGame{}
	}
	return nil, paymentGame
}

func (pg *PaymentGameRepo) DeletePaymentGame(ctx context.Context, input entity.DeletePaymentGameParams) (error, entity.PaymentGame) {
	paymentGame, err := pg.sql.DeletePaymentGame(ctx, input)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("payment not found"), entity.PaymentGame{}
		}
		return err, entity.PaymentGame{}
	}
	return nil, paymentGame
}

func (pg *PaymentGameRepo) GetPaymentGames(ctx context.Context, paymentID string) (error, []string) {
	paymentGames, err := pg.sql.GetPaymentGames(ctx, paymentID)
	if err != nil {
		fmt.Printf(err.Error())
		return errors.New("Cannot get all games"), []string{}
	}
	return nil, paymentGames
}
