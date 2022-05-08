//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
//go:generate goimports -w --local "github.com/karamaru-alpha/melt" mock_$GOPACKAGE/mock_$GOFILE
package auth

import (
	"context"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/karamaru-alpha/melt/pkg/domain/database"
	"github.com/karamaru-alpha/melt/pkg/domain/entity"
	"github.com/karamaru-alpha/melt/pkg/domain/repository"
	"github.com/karamaru-alpha/melt/pkg/merrors"
	"github.com/karamaru-alpha/melt/pkg/util"
)

type Interactor interface {
	Signup(ctx context.Context, name string) (accessToken string, refreshToken string, err error)
	RefreshToken(ctx context.Context, refreshToken string) (accessToken string, err error)
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

func (i *interactor) Signup(ctx context.Context, name string) (accessToken string, refreshToken string, err error) {
	if err := i.txManager.Transaction(ctx, func(ctx context.Context, tx database.Tx) error {
		// User作成
		userID, err := i.ulidGenerator.Generate()
		if err != nil {
			return merrors.Stack(err)
		}
		if err := i.userRepository.Insert(ctx, tx, &entity.User{
			ID:   userID,
			Name: name,
		}); err != nil {
			return merrors.Stack(err)
		}
		// accessToken生成
		accessToken, err = i.generateAccessToken(userID)
		if err != nil {
			return merrors.Stack(err)
		}
		// refreshToken生成
		refreshToken, err = i.generateRefreshToken(userID)
		if err != nil {
			return merrors.Stack(err)
		}
		return nil
	}); err != nil {
		return "", "", merrors.Stack(err)
	}

	return accessToken, refreshToken, nil
}

func (i *interactor) RefreshToken(ctx context.Context, refreshToken string) (accessToken string, err error) {
	if err := i.txManager.Transaction(ctx, func(ctx context.Context, tx database.Tx) error {
		// refreshTokenからuserIDを取り出す
		token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, merrors.Newf(merrors.Unauthenticated, "token method is invalid")
			}
			return []byte(os.Getenv("TOKEN_SIGNED_STRING")), nil
		})
		if err != nil {
			return merrors.Wrap(err, merrors.Unauthenticated)
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return merrors.Newf(merrors.Unauthenticated, "token claim is invalid")
		}
		userID := claims["sub"].(string)
		// 新しいaccessTokenを生成
		accessToken, err = i.generateAccessToken(userID)
		if err != nil {
			return merrors.Stack(err)
		}
		return nil
	}); err != nil {
		return "", merrors.Stack(err)
	}

	return accessToken, nil
}

func (i *interactor) generateAccessToken(userID string) (string, error) {
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}).SignedString([]byte(os.Getenv("TOKEN_SIGNED_STRING")))
	if err != nil {
		return "", merrors.Wrap(err, merrors.Internal)
	}
	return accessToken, nil
}

func (i *interactor) generateRefreshToken(userID string) (string, error) {
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
	}).SignedString([]byte(os.Getenv("TOKEN_SIGNED_STRING")))
	if err != nil {
		return "", merrors.Wrap(err, merrors.Internal)
	}
	return refreshToken, nil
}
