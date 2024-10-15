package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"pixels.com/jwtapp/model"
)

type DBConfig struct {
	DB *gorm.DB
}

var DBconfig DBConfig

func Connect() {

	dsn := "host=localhost port=5432 user=postgres password=password dbname=gouprofdb sslmode=disable"

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail to connect database", err)
	}
	log.Println("Database Connection Established")
	// Migrate models
	DB.AutoMigrate(&model.Article{}, &model.User{})
	DBconfig = DBConfig{DB}
}
