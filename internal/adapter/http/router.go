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
	// employee系
	employeeAPIRoot        = apiVersion + "/employee"
	employeeIDParam        = "employee_id"
	employeeFirstNameParam = "employee_first_name"
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

	// employee系
	employeeRepository := mysql.NewEmployeeRepository()
	employeeService := service.NewEmployeeService(employeeRepository)
	employeeUsecase := usecase.NewEmployeeUsecase(employeeService)
	employeeGroup := e.Group(employeeAPIRoot)
	{
		handler := NewEmployeeHandler(employeeUsecase)

		// v1/employee/{employee_id}
		path := fmt.Sprintf("/:%s", employeeIDParam)
		employeeGroup.GET(path, handler.FindOneByEmpNo())

		// v1/employee/{employee_first_name}
		path = fmt.Sprintf("/:%s", employeeFirstNameParam)
		employeeGroup.GET(path, handler.FindAllByFirstName())
	}

	return e
}
