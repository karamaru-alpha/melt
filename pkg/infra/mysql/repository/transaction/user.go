package transaction

import (
	"context"
	"database/sql"

	"github.com/karamaru-alpha/melt/pkg/domain/database"
	"github.com/karamaru-alpha/melt/pkg/domain/entity/transaction"
	repo "github.com/karamaru-alpha/melt/pkg/domain/repository/transaction"
	"github.com/karamaru-alpha/melt/pkg/infra/mysql"
	model "github.com/karamaru-alpha/melt/pkg/infra/mysql/model/transaction"
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

func (r *userRepository) Insert(ctx context.Context, _tx database.Tx, entity *transaction.User) error {
	tx, err := mysql.ExtractTx(_tx)
	if err != nil {
		return merrors.Stack(err)
	}
	if _, err := tx.ExecContext(ctx, `INSERT INTO user (id, name) VALUES (:id, :name)`, model.NewUser(entity)); err != nil {
		return merrors.Wrap(err, merrors.Internal)
	}
	return nil
}
