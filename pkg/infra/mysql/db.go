package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/karamaru-alpha/melt/pkg/merrors"
)

const (
	defaultMaxIdleConns = 2
	defaultMaxOpenConns = 100
)

func New() *sql.DB {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_ADDR"),
		os.Getenv("MYSQL_DB"),
	)
	
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(merrors.Wrapf(err, merrors.Internal, "Unable to open mysql connection. data source: %s", dataSourceName))
	}
	db.SetMaxIdleConns(defaultMaxIdleConns)
	db.SetMaxOpenConns(defaultMaxOpenConns)
	db.SetConnMaxLifetime(defaultMaxOpenConns * time.Second)

	return db
}
