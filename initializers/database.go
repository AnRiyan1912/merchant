package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"rename.com/andreriyant/go-crud/models"
)

var DB *gorm.DB

func ConnectToDB() {
    dsn := os.Getenv("DB_URL")
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed connecting to database")
    }

    database.AutoMigrate(&models.User{})
    database.AutoMigrate(&models.Person{})
    database.AutoMigrate(&models.Transaction{})

    DB = database

}
