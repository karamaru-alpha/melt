//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
//go:generate goimports -w --local "github.com/karamaru-alpha/melt" mock_$GOPACKAGE/mock_$GOFILE
package user

import (
	"context"

	"github.com/karamaru-alpha/melt/pkg/domain/database"
	"github.com/karamaru-alpha/melt/pkg/domain/entity"
	"github.com/karamaru-alpha/melt/pkg/domain/repository"
	"github.com/karamaru-alpha/melt/pkg/merrors"
	"github.com/karamaru-alpha/melt/pkg/util"
)

type Interactor interface {
	Create(ctx context.Context, name string) error
}

type interactor struct {
	ulidGenerator  util.ULIDGenerator
	userRepository repository.UserRepository
	txManager      database.TxManager
}

func New(ulidGenerator util.ULIDGenerator, userRepository repository.UserRepository, txManager database.TxManager) Interactor {
	return &interactor{
		ulidGenerator,
		userRepository,
		txManager,
	}
}

func (i *interactor) Create(ctx context.Context, name string) error {
	if err := i.txManager.Transaction(ctx, func(ctx context.Context, tx database.Tx) error {
		// ID生成
		id, err := i.ulidGenerator.Generate()
		if err != nil {
			return merrors.Stack(err)
		}
		// User作成
		if err := i.userRepository.Insert(ctx, tx, &entity.User{
			ID:   id,
			Name: name,
		}); err != nil {
			return merrors.Stack(err)
		}
		return nil
	}); err != nil {
		return merrors.Stack(err)
	}
	return nil
}
