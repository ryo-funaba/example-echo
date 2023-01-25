package service

import (
	"context"

	"github.com/ryo-funaba/example_echo/internal/domain/model"
	"github.com/ryo-funaba/example_echo/internal/domain/repository"
	"github.com/ryo-funaba/example_echo/internal/utils/enum"
)

type UserService interface {
	FindOneByID(ctx context.Context, id int) (*model.User, error)
	FindAllByName(ctx context.Context, name string) ([]*model.User, error)
}

type userService struct {
	cluster        repository.Cluster
	userRepository repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{userRepository: r}
}

func (s *userService) FindOneByID(ctx context.Context, id int) (*model.User, error) {
	conn, err := s.cluster.GetBoilCtxExecutor(ctx, enum.Secondary, enum.Main.String())
	if err != nil {
		return nil, err
	}

	return s.userRepository.FindOneByID(ctx, conn, uint(id))
}

func (s *userService) FindAllByName(ctx context.Context, name string) ([]*model.User, error) {
	conn, err := s.cluster.GetBoilCtxExecutor(ctx, enum.Secondary, enum.Main.String())
	if err != nil {
		return nil, err
	}

	return s.userRepository.FindAllByName(ctx, conn, name)
}
