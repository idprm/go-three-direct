package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"waki.mobi/go-yatta-h3i/src/config"
)

func InitDB(cfg *config.Secret) *sql.DB {
	dsn := cfg.Db.Source

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return db
}

func InitGormDB(cfg *config.Secret) *gorm.DB {
	dsn := cfg.Db.ScGorm

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
