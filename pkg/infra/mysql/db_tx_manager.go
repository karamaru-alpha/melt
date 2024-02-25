package mysql

import (
	"context"
	"log"

	"gorm.io/gorm"

	"github.com/karamaru-alpha/melt/pkg/domain/database"
	"github.com/karamaru-alpha/melt/pkg/merrors"
)

type dbTxManager struct {
	db *gorm.DB
}

func NewDBTxManager(db *gorm.DB) database.TxManager {
	return &dbTxManager{db}
}

func (m *dbTxManager) Transaction(ctx context.Context, f func(ctx context.Context, tx database.Tx) error) (err error) {
	tx := m.db.Begin()
	defer func() {
		// ラップした関数内でpanicが起きた場合
		if p := recover(); p != nil {
			e := tx.Rollback().Error
			if e != nil {
				log.Println(merrors.Wrapf(e, merrors.Internal, "fail to MySQL RollBack"))
			}
			panic(p) // re-throw panic after Rollback
		}
		// ラップした関数内でエラーが返された場合
		if err != nil {
			e := tx.Rollback().Error
			if e != nil {
				log.Println(merrors.Wrapf(e, merrors.Internal, "fail to MySQL RollBack"))
			}
			return
		}

		if e := tx.Commit().Error; e != nil {
			log.Println(merrors.Wrapf(e, merrors.Internal, "fail to MySQL RollBack"))
		}
	}()

	// ラップした処理の実行
	if err := f(ctx, &dbTx{tx}); err != nil {
		return merrors.Stack(err)
	}

	return nil
}

type dbTx struct {
	tx *gorm.DB
}

func (t *dbTx) Commit() error {
	return t.tx.Commit().Error
}

func (t *dbTx) Rollback() error {
	return t.tx.Rollback().Error
}

// ExtractTx 抽象Txからsql.Txを取得する
func ExtractTx(tx database.Tx) (*gorm.DB, error) {
	t, ok := tx.(*dbTx)
	if !ok {
		return nil, merrors.Newf(merrors.Internal, "fail to get tx object")
	}
	return t.tx, nil
}
