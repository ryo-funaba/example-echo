package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ryo-funaba/example_echo/internal/usecase"
	"github.com/ryo-funaba/example_echo/internal/utils/errorutil"
)

type employeeHandler struct {
	usecase usecase.EmployeeUsecase
}

func NewEmployeeHandler(u usecase.EmployeeUsecase) *employeeHandler {
	return &employeeHandler{usecase: u}
}

func (h *employeeHandler) FindOneByEmpNo() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		id, err := strconv.Atoi(c.Param(employeeIDParam))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		emp, err := h.usecase.FindOneByEmpNo(ctx, id)
		if err != nil {
			return c.JSON(errorutil.ErrorCode(err), err.Error())
		}

		return c.JSON(http.StatusOK, emp)
	}
}

func (h *employeeHandler) FindAllByFirstName() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		name := c.Param(employeeFirstNameParam)

		emp, err := h.usecase.FindAllByFirstName(ctx, name)
		if err != nil {
			return c.JSON(errorutil.ErrorCode(err), err.Error())
		}

		return c.JSON(http.StatusOK, emp)
	}
}
