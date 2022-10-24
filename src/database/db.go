package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"waki.mobi/go-yatta-h3i/src/config"
	"waki.mobi/go-yatta-h3i/src/pkg/models"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func Connect() {

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

	log.Println("Connected to database successfully")

	// DEBUG ON CONSOLE
	db.Logger = logger.Default.LogMode(logger.Info)

	// TODO: Add migrations
	db.AutoMigrate(
		&models.Config{},
		&models.Adnet{},
		&models.Blacklist{},
		&models.Content{},
		&models.Service{},
		&models.Transaction{},
		&models.Subscription{},
	)

	// TODO: Seed records
	var config []models.Config
	var adnet []models.Adnet
	var content []models.Content
	var service []models.Service

	resultConfig := db.Find(&config)
	resultAdnet := db.Find(&adnet)
	resultContent := db.Find(&content)
	resultService := db.Find(&service)

	if resultConfig.RowsAffected == 0 {
		for i, _ := range configs {
			db.Model(&models.Config{}).Create(&configs[i])
		}
	}

	if resultAdnet.RowsAffected == 0 {
		for i, _ := range adnets {
			db.Model(&models.Adnet{}).Create(&adnets[i])
		}
	}

	if resultContent.RowsAffected == 0 {
		for i, _ := range contents {
			db.Model(&models.Content{}).Create(&contents[i])
		}
	}

	if resultService.RowsAffected == 0 {
		for i, _ := range services {
			db.Model(&models.Service{}).Create(&services[i])
		}
	}

	Database = DbInstance{
		Db: db,
	}
}

var configs = []models.Config{
	{
		Name:  "AUTO_MESSAGE_SENDBIRD",
		Value: "Hi, I'm @v1, please describe the symptoms you are feeling",
	},
	{
		Name:  "AUTO_MESSAGE_SENDBIRD",
		Value: "Hi, Saya @v1 silahkan jelaskan keluhan kamu",
	},
}

var adnets = []models.Adnet{
	{
		Name:  "AUTO_MESSAGE_SENDBIRD",
		Value: "Hi, I'm @v1, please describe the symptoms you are feeling",
	},
	{
		Name:  "AUTO_MESSAGE_SENDBIRD",
		Value: "Hi, Saya @v1 silahkan jelaskan keluhan kamu",
	},
}

var contents = []models.Content{
	{
		Name:  "AUTO_MESSAGE_SENDBIRD",
		Value: "Hi, I'm @v1, please describe the symptoms you are feeling",
	},
	{
		Name:  "AUTO_MESSAGE_SENDBIRD",
		Value: "Hi, Saya @v1 silahkan jelaskan keluhan kamu",
	},
}

var services = []models.Service{
	{
		Name: "AUTO_MESSAGE_SENDBIRD",
	},
	{
		Name: "AUTO_MESSAGE_SENDBIRD",
	},
}
