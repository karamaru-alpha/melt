package user

import (
	"context"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/karamaru-alpha/melt/pkg/domain/database/mock_database"
	"github.com/karamaru-alpha/melt/pkg/domain/entity"
	"github.com/karamaru-alpha/melt/pkg/domain/repository/mock_repository"
	"github.com/karamaru-alpha/melt/pkg/merrors"
	testutil "github.com/karamaru-alpha/melt/pkg/test"
	"github.com/karamaru-alpha/melt/pkg/util/mock_util"
)

type mocks struct {
	userRepository *mock_repository.MockUserRepository
	ulidGenerator  *mock_util.MockULIDGenerator
	tx             *mock_database.MockTx
}

func newWithMocks(t *testing.T) (context.Context, *service, *mocks) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	userRepository := mock_repository.NewMockUserRepository(ctrl)
	ulidGenerator := mock_util.NewMockULIDGenerator(ctrl)
	tx := mock_database.NewMockTx(ctrl)
	return ctx,
		New(userRepository, ulidGenerator).(*service),
		&mocks{userRepository, ulidGenerator, tx}
}

func Test_Service(t *testing.T) {
	for name, tt := range map[string]struct {
		name string
		err  error
		mock func(ctx context.Context, m *mocks)
	}{
		"正常系": {
			name: strings.Repeat("a", 10),
			mock: func(ctx context.Context, m *mocks) {
				name := strings.Repeat("a", 10)
				m.userRepository.EXPECT().SelectByName(ctx, m.tx, name).Return(nil, nil).Times(1)
				m.ulidGenerator.EXPECT().Generate().Return("id", nil).Times(1)
				m.userRepository.EXPECT().Insert(ctx, m.tx, &entity.User{
					ID:   "id",
					Name: name,
				}).Return(nil).Times(1)

			},
		},
		"異常系: 既に存在するname": {
			name: strings.Repeat("a", 10),
			mock: func(ctx context.Context, m *mocks) {
				name := strings.Repeat("a", 10)
				m.userRepository.EXPECT().SelectByName(ctx, m.tx, name).Return([]*entity.User{{}}, nil).Times(1)
			},
			err: merrors.Newf(merrors.InvalidArgument, "user is already exist. name: %s", strings.Repeat("a", 10)),
		},
		"異常系: nameが長すぎる": {
			name: strings.Repeat("a", 11),
			err:  merrors.Newf(merrors.InvalidArgument, "user name len should be %d~%d", 2, 10),
		},
		"異常系: nameが短すぎる": {
			name: strings.Repeat("a", 1),
			err:  merrors.Newf(merrors.InvalidArgument, "user name len should be %d~%d", 2, 10),
		},
	} {
		t.Run(name, func(t *testing.T) {
			ctx, s, m := newWithMocks(t)
			if tt.mock != nil {
				tt.mock(ctx, m)
			}

			err := s.Create(ctx, m.tx, tt.name)
			testutil.EqualMeltError(t, tt.err, err)
		})
	}

}
