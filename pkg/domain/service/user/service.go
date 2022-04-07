//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
//go:generate goimports -w --local "github.com/karamaru-alpha/melt" mock_$GOPACKAGE/mock_$GOFILE
package user

import (
	"context"
	"unicode/utf8"

	"github.com/karamaru-alpha/melt/pkg/domain/database"
	"github.com/karamaru-alpha/melt/pkg/domain/entity"
	"github.com/karamaru-alpha/melt/pkg/domain/repository"
	"github.com/karamaru-alpha/melt/pkg/merrors"
	"github.com/karamaru-alpha/melt/pkg/util"
)

type Service interface {
	Create(ctx context.Context, tx database.Tx, name string) error
}

type service struct {
	userRepository repository.UserRepository
	ulidGenerator  util.ULIDGenerator
}

func New(userRepository repository.UserRepository, ulidGenerator util.ULIDGenerator) Service {
	return &service{
		userRepository: userRepository,
		ulidGenerator:  ulidGenerator,
	}
}

func (s *service) Create(ctx context.Context, tx database.Tx, name string) error {
	// バリデーション
	if utf8.RuneCountInString(name) > entity.UserNameMaxLen || utf8.RuneCountInString(name) < entity.UserNameMinLen {
		return merrors.Newf(merrors.InvalidArgument, "user name len should be %d~%d", entity.UserNameMinLen, entity.UserNameMaxLen)
	}

	// 重複確認
	users, err := s.userRepository.SelectByName(ctx, tx, name)
	if err != nil {
		return merrors.Stack(err)
	}
	if len(users) != 0 {
		return merrors.Newf(merrors.InvalidArgument, "user is already exist. name: %s", name)
	}

	// ID生成
	id, err := s.ulidGenerator.Generate()
	if err != nil {
		return merrors.Stack(err)
	}

	// 生成
	if err := s.userRepository.Insert(ctx, tx, &entity.User{
		ID:   id,
		Name: name,
	}); err != nil {
		return merrors.Stack(err)
	}
	return nil
}
