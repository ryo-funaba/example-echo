package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ryo-funaba/example_echo/internal/usecase"
	"github.com/ryo-funaba/example_echo/internal/utils/errorutil"
)

type userHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) *userHandler {
	return &userHandler{usecase: u}
}

func (h *userHandler) FindOneByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		id, err := strconv.Atoi(c.Param(userIDParam))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		user, err := h.usecase.FindOneByID(ctx, id)
		if err != nil {
			return c.JSON(errorutil.ErrorCode(err), err.Error())
		}

		return c.JSON(http.StatusOK, user)
	}
}

func (h *userHandler) FindAllByName() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		name := c.Param(userNameParam)

		users, err := h.usecase.FindAllByName(ctx, name)
		if err != nil {
			return c.JSON(errorutil.ErrorCode(err), err.Error())
		}

		return c.JSON(http.StatusOK, users)
	}
}
