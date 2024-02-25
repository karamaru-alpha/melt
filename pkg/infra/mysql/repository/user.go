package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/karamaru-alpha/melt/pkg/domain/database"
	domain "github.com/karamaru-alpha/melt/pkg/domain/entity"
	repo "github.com/karamaru-alpha/melt/pkg/domain/repository"
	"github.com/karamaru-alpha/melt/pkg/infra/mysql"
	"github.com/karamaru-alpha/melt/pkg/infra/mysql/model"
	"github.com/karamaru-alpha/melt/pkg/merrors"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repo.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) SelectByName(ctx context.Context, name string) ([]*domain.User, error) {
	var users []*model.User
	if err := r.db.WithContext(ctx).Where("name = ?", name).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, merrors.Wrap(err, merrors.Internal)
	}
	entities := make([]*domain.User, len(users))
	for _, v := range users {
		entities = append(entities, v.ToEntity())
	}
	return entities, nil
}

func (r *userRepository) Insert(ctx context.Context, _tx database.Tx, entity *domain.User) error {
	tx, err := mysql.ExtractTx(_tx)
	if err != nil {
		return merrors.Stack(err)
	}
	if err := tx.WithContext(ctx).Create(model.NewUser(entity)).Error; err != nil {
		return merrors.Wrap(err, merrors.Internal)
	}
	return nil
}
