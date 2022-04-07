package repository

import (
	"context"
	"database/sql"

	"github.com/karamaru-alpha/melt/pkg/domain/database"
	domain "github.com/karamaru-alpha/melt/pkg/domain/entity"
	repo "github.com/karamaru-alpha/melt/pkg/domain/repository"
	"github.com/karamaru-alpha/melt/pkg/infra/mysql"
	"github.com/karamaru-alpha/melt/pkg/infra/mysql/model"
	"github.com/karamaru-alpha/melt/pkg/merrors"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repo.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Insert(ctx context.Context, _tx database.Tx, entity *domain.User) error {
	tx, err := mysql.ExtractTx(_tx)
	if err != nil {
		return merrors.Stack(err)
	}
	if _, err := tx.ExecContext(ctx, `INSERT INTO user (id, name) VALUES (:id, :name)`, model.NewUser(entity)); err != nil {
		return merrors.Wrap(err, merrors.Internal)
	}
	return nil
}

func (r *userRepository) SelectByName(ctx context.Context, _tx database.Tx, name string) ([]*domain.User, error) {
	// TODO
	//tx, err := mysql.ExtractTx(_tx)
	//if err != nil {
	//	return nil, merrors.Stack(err)
	//}
	//
	//rows, err := tx.QueryContext(ctx, `SELECT * FROM user WHERE name = ?`, name)
	//if err != nil {
	//	return nil, merrors.Wrap(err, merrors.Internal)
	//}
	return nil, nil
}
