package http

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ryo-funaba/example_echo/internal/domain/service"
	"github.com/ryo-funaba/example_echo/internal/infra/mysql"
	"github.com/ryo-funaba/example_echo/internal/usecase"
)

const (
	apiVersion = "/v1"
	// ヘルスチェック系
	healthCheckRoot = "/health_check"
	// user系
	userAPIRoot   = apiVersion + "/user"
	userIDParam   = "user_id"
	userNameParam = "user_name"
)

func InitRouter() *echo.Echo {
	e := echo.New()
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	// ヘルスチェック系
	healthCheckGroup := e.Group(healthCheckRoot)
	{
		path := ""
		healthCheckGroup.GET(path, healthCheck)
	}

	// user系
	userRepository := mysql.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userUsecase := usecase.NewUserUsecase(userService)
	userGroup := e.Group(userAPIRoot)
	{
		handler := NewUserHandler(userUsecase)

		// v1/user/{user_id}
		path := fmt.Sprintf("/:%s", userIDParam)
		userGroup.GET(path, handler.FindOneByID())

		// v1/user/{user_name}
		path = fmt.Sprintf("/:%s", userNameParam)
		userGroup.GET(path, handler.FindAllByName())
	}

	return e
}
