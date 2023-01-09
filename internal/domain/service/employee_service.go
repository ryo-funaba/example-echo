package service

import (
	"context"

	"github.com/ryo-funaba/example_echo/internal/domain/model"
	"github.com/ryo-funaba/example_echo/internal/domain/repository"
	"github.com/ryo-funaba/example_echo/internal/utils/enum"
)

type EmployeeService interface {
	FindOneByEmpNo(ctx context.Context, empNo int) (*model.Employee, error)
	FindAllByFirstName(ctx context.Context, firstName string) ([]*model.Employee, error)
}

type employeeService struct {
	cluster            repository.Cluster
	employeeRepository repository.EmployeeRepository
}

func NewEmployeeService(r repository.EmployeeRepository) EmployeeService {
	return &employeeService{employeeRepository: r}
}

func (s *employeeService) FindOneByEmpNo(ctx context.Context, empNo int) (*model.Employee, error) {
	conn, err := s.cluster.GetBoilCtxExecutor(ctx, enum.Secondary, enum.Main.String())
	if err != nil {
		return nil, err
	}

	return s.employeeRepository.FindOneByEmpNo(ctx, conn, empNo)
}

func (s *employeeService) FindAllByFirstName(ctx context.Context, firstName string) ([]*model.Employee, error) {
	conn, err := s.cluster.GetBoilCtxExecutor(ctx, enum.Secondary, enum.Main.String())
	if err != nil {
		return nil, err
	}

	return s.employeeRepository.FindAllByFirstName(ctx, conn, firstName)
}
