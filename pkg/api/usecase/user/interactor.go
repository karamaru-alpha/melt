//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
//go:generate goimports -w --local "github.com/karamaru-alpha/melt" mock_$GOPACKAGE/mock_$GOFILE
package user

import (
	"context"

	"github.com/karamaru-alpha/melt/pkg/domain/database"
	"github.com/karamaru-alpha/melt/pkg/domain/service/user"
	"github.com/karamaru-alpha/melt/pkg/merrors"
)

type Interactor interface {
	Create(ctx context.Context, name string) error
}

type interactor struct {
	userService user.Service
	txManager   database.TxManager
}

func New(userService user.Service, txManager database.TxManager) Interactor {
	return &interactor{
		userService: userService,
		txManager:   txManager,
	}
}

func (i *interactor) Create(ctx context.Context, name string) error {
	if err := i.txManager.Transaction(ctx, func(ctx context.Context, tx database.Tx) error {
		if err := i.userService.Create(ctx, tx, name); err != nil {
			return merrors.Stack(err)
		}
		return nil
	}); err != nil {
		return merrors.Stack(err)
	}
	return nil
}
