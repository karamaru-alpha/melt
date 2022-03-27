package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/karamaru-alpha/melt/pkg/merrors"
)

const (
	defaultMaxIdleConns = 2
	defaultMaxOpenConns = 100
)

type Config struct {
	Addr     string
	User     string
	Password string
	DB       string
}

func New(c *Config) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local", c.User, c.Password, c.Addr, c.DB)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, merrors.Wrapf(err, merrors.Internal, "Unable to open mysql connection. data source: %s", dataSourceName)
	}
	db.SetMaxIdleConns(defaultMaxIdleConns)
	db.SetMaxOpenConns(defaultMaxOpenConns)
	db.SetConnMaxLifetime(defaultMaxOpenConns * time.Second)
	return db, nil
}
