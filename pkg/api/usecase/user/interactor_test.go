package user

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/karamaru-alpha/melt/pkg/domain/database/mock_database"
	"github.com/karamaru-alpha/melt/pkg/domain/entity"
	"github.com/karamaru-alpha/melt/pkg/domain/repository/mock_repository"
	"github.com/karamaru-alpha/melt/pkg/domain/service/user/mock_user"
	"github.com/karamaru-alpha/melt/pkg/util/mock_util"
)

type mocks struct {
	userService    *mock_user.MockService
	userRepository *mock_repository.MockUserRepository
	ulidGenerator  *mock_util.MockULIDGenerator
	txManager      *mock_database.MockTxManager
	tx             *mock_database.MockTx
}

func newWithMocks(t *testing.T) (context.Context, *interactor, *mocks) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	userService := mock_user.NewMockService(ctrl)
	userRepository := mock_repository.NewMockUserRepository(ctrl)
	ulidGenerator := mock_util.NewMockULIDGenerator(ctrl)
	txManager := mock_database.NewMockTxManager(ctrl)
	tx := mock_database.NewMockTx(ctrl)
	return ctx,
		New(userService, userRepository, ulidGenerator, txManager).(*interactor),
		&mocks{userService, userRepository, ulidGenerator, txManager, tx}
}

func Test_Create(t *testing.T) {
	ctx, i, m := newWithMocks(t)
	m.txManager.EXPECTTransaction(ctx, m.tx, 1)
	name := "name"
	id := "id"
	m.userService.EXPECT().ValidateUserName(ctx, m.tx, name).Return(nil).Times(1)
	m.ulidGenerator.EXPECT().Generate().Return(id, nil).Times(1)
	m.userRepository.EXPECT().Insert(ctx, m.tx, &entity.User{
		ID:   id,
		Name: name,
	}).Return(nil).Times(1)

	err := i.Create(ctx, name)
	assert.NoError(t, err)
}
