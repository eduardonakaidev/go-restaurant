package connection

import (
	"fmt"
	"log"
	"os"

	"github.com/eduardonakaidev/go-restaurant/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectPostgresDB() {
	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL"),
	)
	DB, err = gorm.Open(postgres.Open(url))
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	fmt.Println("Database connected successfully !")

	DB.AutoMigrate(
		&models.Product{},
	)
}
