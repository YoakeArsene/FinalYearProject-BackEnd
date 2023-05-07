package http

import (
	sql "GDN-delivery-management/db/sql"
	repo "GDN-delivery-management/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PaymentHandler struct {
	PaymentRepo repo.IPaymentRepo
}

type CreatePaymentRequest struct {
	ID     string  `json:"id"`
	UserID string  `json:"user_id"`
	Price  float64 `json:"price"`
}

func (p *PaymentHandler) CreatePayment(c echo.Context) error {
	req := CreatePaymentRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	payId, err := uuid.NewUUID()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	req.ID = payId.String()
	param := sql.CreatePaymentParams{
		ID:     req.ID,
		UserID: req.UserID,
		Price:  req.Price,
	}
	err, payment := p.PaymentRepo.CreatePayment(c.Request().Context(), param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       payment,
	})
}

type GetUserPaymentsRequest struct {
	UserID string `json:"user_id"`
}

func (p *PaymentHandler) GetUserPayments(c echo.Context) error {
	req := GetUserPaymentsRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err, payment := p.PaymentRepo.GetUserPayments(c.Request().Context(), req.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, payment)
}

type DeletePaymentRequest struct {
	ID string `json:"id"`
}

func (p *PaymentHandler) DeletePayment(c echo.Context) error {
	req := DeletePaymentRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err, payment := p.PaymentRepo.DeletePayment(c.Request().Context(), req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       payment,
	})
}
