package mysql

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	defaultMaxIdleConns = 2
	defaultMaxOpenConns = 100
)

func New() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	) + "?parseTime=true&collation=utf8mb4_bin"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	for i := 0; i < 60; i++ {
		if err == nil {
			break
		}
		log.Printf("waiting db (%ds)...\n", i)
		time.Sleep(time.Second)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("fail to connect db")
	}
	sqlDB.SetMaxIdleConns(defaultMaxIdleConns)
	sqlDB.SetMaxOpenConns(defaultMaxOpenConns)
	sqlDB.SetConnMaxLifetime(defaultMaxOpenConns * time.Second)

	if os.Getenv("ENV") == "local" {
		db.Logger = db.Logger.LogMode(logger.Info)
	}
	return db
}
