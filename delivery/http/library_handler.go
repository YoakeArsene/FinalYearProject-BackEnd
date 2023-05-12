package http

import (
	sql "GDN-delivery-management/db/sql"
	repo "GDN-delivery-management/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type LibraryHandler struct {
	LibraryRepo repo.ILibraryRepo
}

type CreateLibraryRequest struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	GameID int32  `json:"game_id"`
}

func (l *LibraryHandler) CreateLibrary(c echo.Context) error {
	req := CreateLibraryRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	libId, err := uuid.NewUUID()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	req.ID = libId.String()
	param := sql.CreateLibraryParams{
		ID:     req.ID,
		UserID: req.UserID,
		GameID: req.GameID,
	}
	err, library := l.LibraryRepo.CreateLibrary(c.Request().Context(), param)
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
		Data:       library,
	})
}

type GetUserLibraryRequest struct {
	UserID string `json:"user_id"`
}

func (l *LibraryHandler) GetUserLibrary(c echo.Context) error {
	req := GetUserLibraryRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	userId := req.UserID
	userId = c.Param("userid")
	err, library := l.LibraryRepo.GetUserLibrary(c.Request().Context(), userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, library)
}

type DeleteLibraryRequest struct {
	UserID string `json:"user_id"`
	GameID int32  `json:"game_id"`
}

func (l *LibraryHandler) DeleteLibrary(c echo.Context) error {
	req := DeleteLibraryRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	param := sql.DeleteLibraryParams{
		UserID: req.UserID,
		GameID: req.GameID,
	}
	err, library := l.LibraryRepo.DeleteLibrary(c.Request().Context(), param)
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
		Data:       library,
	})
}

type CheckGameInLibraryRequest struct {
	UserID string `json:"user_id"`
	GameID int32  `json:"game_id"`
}

func (l *LibraryHandler) CheckGameInLibrary(c echo.Context) error {
	req := CheckGameInLibraryRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	param := sql.CheckGameInLibraryParams{
		UserID: req.UserID,
		GameID: req.GameID,
	}
	err, library := l.LibraryRepo.CheckGameInLibrary(c.Request().Context(), param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, library)
}
