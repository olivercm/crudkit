package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type Config struct {
	DSN          string
	MaxIdleConns int
	MaxOpenConns int
	MaxLifetime  time.Duration
}

func Open(conf *Config) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", conf.DSN)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	db.DB().SetMaxIdleConns(conf.MaxIdleConns)
	db.DB().SetMaxOpenConns(conf.MaxOpenConns)
	db.DB().SetConnMaxLifetime(conf.MaxLifetime)
	return db, nil
}

func Master(user, password, address string) (*gorm.DB, error) {
	config := &Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s)/liangyin?charset=utf8mb4&parseTime=True&loc=Local",
			user, password, address),
		MaxIdleConns: 16,
		MaxOpenConns: 32,
		MaxLifetime:  3 * time.Hour,
	}
	return Open(config)
}
