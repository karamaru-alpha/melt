package user

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/karamaru-alpha/melt/pkg/domain/database/mock_database"
	"github.com/karamaru-alpha/melt/pkg/domain/service/user/mock_user"
)

type mocks struct {
	userService *mock_user.MockService
	txManager   *mock_database.MockTxManager
	tx          *mock_database.MockTx
}

func newWithMocks(t *testing.T) (context.Context, *interactor, *mocks) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	userService := mock_user.NewMockService(ctrl)
	txManager := mock_database.NewMockTxManager(ctrl)
	tx := mock_database.NewMockTx(ctrl)
	return ctx,
		New(userService, txManager).(*interactor),
		&mocks{userService, txManager, tx}
}

func Test_Create(t *testing.T) {
	ctx, i, m := newWithMocks(t)
	m.txManager.EXPECTTransaction(ctx, m.tx, 1)
	name := "name"
	m.userService.EXPECT().Create(ctx, m.tx, name).Return(nil).Times(1)

	err := i.Create(ctx, name)
	assert.NoError(t, err)
}
