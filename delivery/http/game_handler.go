package http

import (
	sql "GDN-delivery-management/db/sql"
	repo "GDN-delivery-management/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GameHandler struct {
	GameRepo repo.IGameRepo
}

type CreateGameRequest struct {
	ID         string   `json:"id"`
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
	Rating     int32    `json:"rating"`
	Cover      string   `json:"cover"`
	Footage    []string `json:"footage"`
}

func (g *GameHandler) CreateGame(c echo.Context) error {
	req := CreateGameRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	param := sql.CreateGameParams{
		ID:         req.ID,
		Name:       req.Name,
		Surname:    req.Surname,
		Price:      req.Price,
		Desc:       req.Desc,
		Link:       req.Link,
		Release:    req.Release,
		Platforms:  req.Platforms,
		Genre:      req.Genre,
		Developers: req.Developers,
		Publishers: req.Publishers,
		InCart:     false,
		Selected:   false,
		IsHovered:  false,
		IsLiked:    false,
		Rating:     req.Rating,
		Cover:      req.Cover,
		Column18:   req.Footage,
	}
	err, game := g.GameRepo.CreateGame(c.Request().Context(), param)
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
		Data:       game,
	})
}

func (g *GameHandler) GetAllGames(c echo.Context) error {
	err, games := g.GameRepo.GetAllGames(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, games)
}

type UpdateGameRequest struct {
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
	Rating     int32    `json:"rating"`
	Cover      string   `json:"cover"`
	Footage    []string `json:"footage"`
	ID         string   `json:"id"`
}

func (g *GameHandler) UpdateGame(c echo.Context) error {
	req := UpdateGameRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	param := sql.UpdateGameParams{
		Name:       req.Name,
		Surname:    req.Surname,
		Price:      req.Price,
		Desc:       req.Desc,
		Link:       req.Link,
		Release:    req.Release,
		Platforms:  req.Platforms,
		Genre:      req.Genre,
		Developers: req.Developers,
		Publishers: req.Publishers,
		InCart:     false,
		Selected:   false,
		IsHovered:  false,
		IsLiked:    false,
		Rating:     req.Rating,
		Cover:      req.Cover,
		Column17:   req.Footage,
		ID:         req.ID,
	}
	err, game := g.GameRepo.UpdateGame(c.Request().Context(), param)
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
		Data:       game,
	})
}

type DeleteGameRequest struct {
	ID string `json:"id"`
}

func (g *GameHandler) DeleteGame(c echo.Context) error {
	req := DeleteGameRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	err, game := g.GameRepo.DeleteGame(c.Request().Context(), req.ID)
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
		Data:       game,
	})
}
