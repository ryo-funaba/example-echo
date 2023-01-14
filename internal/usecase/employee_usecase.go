package usecase

import (
	"context"

	"github.com/ryo-funaba/example_echo/internal/domain/service"
	"github.com/ryo-funaba/example_echo/internal/usecase/model"
)

type EmployeeUsecase interface {
	FindOneByEmpNo(ctx context.Context, empNo int) (*model.Employee, error)
	FindAllByFirstName(ctx context.Context, firstName string) ([]*model.Employee, error)
}

type employeeUsecase struct {
	employeeService service.EmployeeService
}

func NewEmployeeUsecase(s service.EmployeeService) EmployeeUsecase {
	return &employeeUsecase{employeeService: s}
}

func (u *employeeUsecase) FindOneByEmpNo(ctx context.Context, empNo int) (*model.Employee, error) {
	e, err := u.employeeService.FindOneByEmpNo(ctx, empNo)
	if err != nil {
		return nil, err
	}

	return model.EmployeeFromDomainModel(e), nil
}

func (u *employeeUsecase) FindAllByFirstName(ctx context.Context, firstName string) ([]*model.Employee, error) {
	e, err := u.employeeService.FindAllByFirstName(ctx, firstName)
	if err != nil {
		return nil, err
	}

	s := make([]*model.Employee, 0, len(e))

	for _, v := range e {
		s = append(s, model.EmployeeFromDomainModel(v))
	}

	return s, nil
}
