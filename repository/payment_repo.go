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

type IPaymentRepo interface {
	CreatePayment(ctx context.Context, input entity.CreatePaymentParams) (error, entity.Payment)
	DeletePayment(ctx context.Context, id string) (error, entity.Payment)
	GetUserPayments(ctx context.Context, userID string) (error, []entity.Payment)
}

type PaymentRepo struct {
	sql *entity.Queries
}

func NewPaymentRepo(sql *entity.Queries) IPaymentRepo {
	return &PaymentRepo{
		sql: sql,
	}
}

func (p *PaymentRepo) CreatePayment(ctx context.Context, input entity.CreatePaymentParams) (error, entity.Payment) {
	payment, err := p.sql.CreatePayment(ctx, input)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return errors.New("Payment already existed"), entity.Payment{}
			}
		}
		log.Println("err: ", err)
		return errors.New("Failed to add payment"), entity.Payment{}
	}
	return nil, payment
}

func (p *PaymentRepo) DeletePayment(ctx context.Context, id string) (error, entity.Payment) {
	payment, err := p.sql.DeletePayment(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("payment not found"), entity.Payment{}
		}
		return err, entity.Payment{}
	}
	return nil, payment
}

func (p *PaymentRepo) GetUserPayments(ctx context.Context, userID string) (error, []entity.Payment) {
	payment, err := p.sql.GetUserPayment(ctx, userID)
	if err != nil {
		fmt.Printf(err.Error())
		return errors.New("Cannot get all games"), []entity.Payment{}
	}
	return nil, payment
}
