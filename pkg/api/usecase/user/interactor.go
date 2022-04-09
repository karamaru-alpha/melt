//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
//go:generate goimports -w --local "github.com/karamaru-alpha/melt" mock_$GOPACKAGE/mock_$GOFILE
package user

import (
	"context"

	"github.com/karamaru-alpha/melt/pkg/domain/database"
	"github.com/karamaru-alpha/melt/pkg/domain/entity"
	"github.com/karamaru-alpha/melt/pkg/domain/repository"
	"github.com/karamaru-alpha/melt/pkg/domain/service/user"
	"github.com/karamaru-alpha/melt/pkg/merrors"
	"github.com/karamaru-alpha/melt/pkg/util"
)

type Interactor interface {
	Create(ctx context.Context, name string) error
}

type interactor struct {
	userService    user.Service
	userRepository repository.UserRepository
	ulidGenerator  util.ULIDGenerator
	txManager      database.TxManager
}

func New(userService user.Service, userRepository repository.UserRepository, ulidGenerator util.ULIDGenerator, txManager database.TxManager) Interactor {
	return &interactor{
		userService:    userService,
		userRepository: userRepository,
		ulidGenerator:  ulidGenerator,
		txManager:      txManager,
	}
}

func (i *interactor) Create(ctx context.Context, name string) error {
	if err := i.txManager.Transaction(ctx, func(ctx context.Context, tx database.Tx) error {
		// 名前のバリデーション
		if err := i.userService.ValidateUserName(ctx, tx, name); err != nil {
			return merrors.Stack(err)
		}
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
