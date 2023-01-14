package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type healthCheckResponse struct {
	Message string `json:"message"`
}

func healthCheck(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		healthCheckResponse{
			Message: "server is healthy",
		},
	)
}
