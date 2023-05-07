package http

import (
	sql "GDN-delivery-management/db/sql"
	repo "GDN-delivery-management/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PaymentGameHandler struct {
	PaymentGameRepo repo.IPaymentGameRepo
}

type CreatePaymentGameRequest struct {
	PaymentID string `json:"payment_id"`
	GameID    int32  `json:"game_id"`
}

func (pg *PaymentGameHandler) CreatePaymentGame(c echo.Context) error {
	req := CreatePaymentGameRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	param := sql.CreatePaymentGameParams{
		PaymentID: req.PaymentID,
		GameID:    req.GameID,
	}
	err, paymentGame := pg.PaymentGameRepo.CreatePaymentGame(c.Request().Context(), param)
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
		Data:       paymentGame,
	})
}

type GetPaymentGamesRequest struct {
	PaymentId string `json:"payment_id"`
}

func (pg *PaymentGameHandler) GetPaymentGames(c echo.Context) error {
	req := GetPaymentGamesRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err, paymentGame := pg.PaymentGameRepo.GetPaymentGames(c.Request().Context(), req.PaymentId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, paymentGame)
}

type DeletePaymentGameRequest struct {
	PaymentID string `json:"payment_id"`
	GameID    int32  `json:"game_id"`
}

func (pg *PaymentGameHandler) DeletePaymentGame(c echo.Context) error {
	req := DeletePaymentGameRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	param := sql.DeletePaymentGameParams{
		PaymentID: req.PaymentID,
		GameID:    req.GameID,
	}
	err, paymentGame := pg.PaymentGameRepo.DeletePaymentGame(c.Request().Context(), param)
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
		Data:       paymentGame,
	})
}
