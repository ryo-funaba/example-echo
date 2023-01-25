package usecase

import (
	"context"

	"github.com/ryo-funaba/example_echo/internal/domain/service"
	"github.com/ryo-funaba/example_echo/internal/usecase/model"
)

type UserUsecase interface {
	FindOneByID(ctx context.Context, id int) (*model.User, error)
	FindAllByName(ctx context.Context, name string) ([]*model.User, error)
}

type userUsecase struct {
	userService service.UserService
}

func NewUserUsecase(s service.UserService) UserUsecase {
	return &userUsecase{userService: s}
}

func (u *userUsecase) FindOneByID(ctx context.Context, id int) (*model.User, error) {
	e, err := u.userService.FindOneByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return model.UserFromDomainModel(e), nil
}

func (u *userUsecase) FindAllByName(ctx context.Context, name string) ([]*model.User, error) {
	e, err := u.userService.FindAllByName(ctx, name)
	if err != nil {
		return nil, err
	}

	s := make([]*model.User, 0, len(e))

	for _, v := range e {
		s = append(s, model.UserFromDomainModel(v))
	}

	return s, nil
}
