package user

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/karamaru-alpha/melt/pkg/api/usecase/user/mock_user"
	pb "github.com/karamaru-alpha/melt/pkg/domain/proto/api"
)

type mocks struct {
	userInteractor *mock_user.MockInteractor
}

func newWithMocks(t *testing.T) (context.Context, *handler, *mocks) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	userInteractor := mock_user.NewMockInteractor(ctrl)
	return ctx,
		New(userInteractor).(*handler),
		&mocks{userInteractor}
}

func Test_Create(t *testing.T) {
	ctx, h, m := newWithMocks(t)
	name := "name"
	m.userInteractor.EXPECT().Create(ctx, name).Return(nil).Times(1)

	_, err := h.Create(ctx, &pb.UserCreateRequest{Name: name})
	assert.NoError(t, err)
}
