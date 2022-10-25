package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"waki.mobi/go-yatta-h3i/src/pkg/config"
	"waki.mobi/go-yatta-h3i/src/pkg/model"
)

var Datasource *NewDatasource

type NewDatasource struct {
	db    *gorm.DB
	sqlDb *sql.DB
}

func (d NewDatasource) DB() *gorm.DB {
	return d.db
}

func (d NewDatasource) SqlDB() *sql.DB {
	return d.sqlDb
}

func Connect() {

	var db *gorm.DB
	var sqlDb *sql.DB

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.ViperEnv("DB_USER"),
		config.ViperEnv("DB_PASS"),
		config.ViperEnv("DB_HOST"),
		config.ViperEnv("DB_PORT"),
		config.ViperEnv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		panic("Could not connect with the database!")
	}

	sqlDb, _ = db.DB()
	sqlDb.SetConnMaxLifetime(time.Minute * 2)
	sqlDb.SetMaxOpenConns(10000)
	sqlDb.SetMaxIdleConns(10000)

	// try to establish connection
	if sqlDb != nil {
		err = sqlDb.Ping()
		if err != nil {
			log.Fatal("cannot connect to db:", err.Error())
		}
	}

	log.Println("Connected to database successfully")

	// DEBUG ON CONSOLE
	db.Logger = logger.Default.LogMode(logger.Info)

	// TODO: Add migrations
	db.AutoMigrate(
		&model.Config{},
		&model.Schedule{},
		&model.Adnet{},
		&model.Keyword{},
		&model.Blacklist{},
		&model.Content{},
		&model.Service{},
		&model.Transaction{},
		&model.Subscription{},
	)

	// TODO: Seed records
	var config []model.Config
	var schedule []model.Schedule
	var content []model.Content
	var service []model.Service
	var keyword []model.Keyword

	resultConfig := db.Find(&config)
	resultSchedule := db.Find(&schedule)
	resultContent := db.Find(&content)
	resultService := db.Find(&service)
	resultKeyword := db.Find(&keyword)

	if resultConfig.RowsAffected == 0 {
		for i, _ := range configs {
			db.Model(&model.Config{}).Create(&configs[i])
		}
	}

	if resultSchedule.RowsAffected == 0 {
		for i, _ := range schedules {
			db.Model(&model.Schedule{}).Create(&schedules[i])
		}
	}

	if resultContent.RowsAffected == 0 {
		for i, _ := range contents {
			db.Model(&model.Content{}).Create(&contents[i])
		}
	}

	if resultService.RowsAffected == 0 {
		for i, _ := range services {
			db.Model(&model.Service{}).Create(&services[i])
		}
	}

	if resultKeyword.RowsAffected == 0 {
		for i, _ := range keywords {
			db.Model(&model.Keyword{}).Create(&keywords[i])
		}
	}

	Datasource = &NewDatasource{db: db, sqlDb: sqlDb}
}
