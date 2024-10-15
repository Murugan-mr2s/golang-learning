package repo

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"pixels.com/userappdb/model"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost port=5432 user=postgres password=password dbname=gouprofdb sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database")
	}
	fmt.Println("Database connection established")
	db.AutoMigrate(&model.User{})
	return db
}
