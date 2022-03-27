//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
//go:generate goimports -w --local "github.com/karamaru-alpha/melt" mock_$GOPACKAGE/mock_$GOFILE
package database

import (
	"context"
)

// TxManager トランザクションマネージャー
type TxManager interface {
	Transaction(ctx context.Context, f func(ctx context.Context, tx Tx) error) error
}

// Tx トランザクション
type Tx interface {
	Commit() error
	Rollback() error
}
