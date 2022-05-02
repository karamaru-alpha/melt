//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
//go:generate goimports -w --local "github.com/karamaru-alpha/melt" mock_$GOPACKAGE/mock_$GOFILE
package repository

import (
	"context"

	"github.com/karamaru-alpha/melt/pkg/domain/database"
	"github.com/karamaru-alpha/melt/pkg/domain/entity"
)

type UserRepository interface {
	SelectByName(ctx context.Context, name string) ([]*entity.User, error)

	Insert(ctx context.Context, tx database.Tx, entity *entity.User) error
}
