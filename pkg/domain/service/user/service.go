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
)

type Service interface {
	ValidateUserName(ctx context.Context, tx database.Tx, name string) error
}

type service struct {
	userRepository repository.UserRepository
}

func New(userRepository repository.UserRepository) Service {
	return &service{
		userRepository: userRepository,
	}
}

func (s *service) ValidateUserName(ctx context.Context, tx database.Tx, name string) error {
	// 文字列のバリデーション
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

	return nil
}
