package repository

import (
	"context"

	"github.com/ryo-funaba/example_echo/internal/domain/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserRepository interface {
	FindOneByID(ctx context.Context, conn boil.ContextExecutor, id uint) (*model.User, error)
	FindAllByName(ctx context.Context, conn boil.ContextExecutor, name string) ([]*model.User, error)
}
