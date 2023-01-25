package mysql

import (
	"context"

	"github.com/ryo-funaba/example_echo/internal/domain/model"
	"github.com/ryo-funaba/example_echo/internal/domain/repository"
	"github.com/ryo-funaba/example_echo/internal/utils/errorutil"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userRepository struct{}

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindOneByID(ctx context.Context, conn boil.ContextExecutor, id uint) (*model.User, error) {
	q := []qm.QueryMod{
		model.UserWhere.ID.EQ(id),
	}

	row, err := model.Users(q...).One(ctx, conn)
	if err != nil {
		return nil, errorutil.Errorf(errorutil.Unknown, err.Error())
	}

	return row, nil
}

func (r *userRepository) FindAllByName(ctx context.Context, conn boil.ContextExecutor, name string) ([]*model.User, error) {
	q := []qm.QueryMod{
		model.UserWhere.Name.EQ(name),
	}

	rows, err := model.Users(q...).All(ctx, conn)
	if err != nil {
		return nil, errorutil.Errorf(errorutil.Unknown, err.Error())
	}

	return rows, nil
}
