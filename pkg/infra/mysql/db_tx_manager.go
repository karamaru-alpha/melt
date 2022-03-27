package mysql

import (
	"context"
	"database/sql"
	"errors"

	"go.uber.org/zap"

	"github.com/karamaru-alpha/melt/pkg/domain/database"
	"github.com/karamaru-alpha/melt/pkg/merrors"
)

type dbTxManager struct {
	db *sql.DB
}

func New(db *sql.DB) database.TxManager {
	return &dbTxManager{db}
}

func (m *dbTxManager) Transaction(ctx context.Context, f func(ctx context.Context, tx database.Tx) error) error {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return merrors.Wrapf(err, merrors.Internal, "fail to begin tx")
	}
	defer func() {
		// ラップした関数内でpanicが起きた場合
		if p := recover(); p != nil {
			e := tx.Rollback()
			if e != nil {
				zap.Error(errors.New("fail to MySQL RollBack"))
			}
			panic(p) // re-throw panic after Rollback
		}
		// ラップした関数内でエラーが返された場合
		if err != nil {
			e := tx.Rollback()
			if e != nil {
				zap.Error(errors.New("fail to MySQL RollBack"))
			}
			return
		}

		if e := tx.Commit(); e != nil {
			zap.Error(errors.New("fail to MySQL RollBack"))
		}
	}()

	// ラップした処理の実行
	if err = f(ctx, &dbTx{tx}); err != nil {
		return merrors.Stack(err)
	}

	return nil
}

type dbTx struct {
	tx *sql.Tx
}

func (t *dbTx) Commit() error {
	return t.tx.Commit()
}

func (t *dbTx) Rollback() error {
	return t.tx.Rollback()
}

// ExtractTx 抽象Txからsql.Txを取得する
func ExtractTx(tx database.Tx) (*sql.Tx, error) {
	t, ok := tx.(*dbTx)
	if !ok {
		return nil, merrors.Newf(merrors.Internal, "fail to get tx object")
	}
	return t.tx, nil
}
