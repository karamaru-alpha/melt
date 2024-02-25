package mock_database //nolint:golint,stylecheck // mockに実装を追加するためにアンダーバーを含むパッケージ名にする必要がある

import (
	"context"

	"github.com/golang/mock/gomock"

	"github.com/karamaru-alpha/melt/pkg/domain/database"
)

func (m *MockTxManager) EXPECTTransaction(ctx context.Context, tx database.Tx, times int) {
	m.EXPECT().Transaction(ctx, gomock.Any()).Times(times).DoAndReturn(
		func(ctx context.Context, f func(context.Context, database.Tx) error) error {
			return f(ctx, tx)
		})
}
