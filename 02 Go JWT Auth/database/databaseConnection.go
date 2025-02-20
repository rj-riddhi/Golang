package database

import (
	"log"

	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Loadenv() (err error) {
	err = nil
	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")

	}
	return err
}
func Connect() *gorm.DB {
	// Load environment variables from .env file
	Loadenv()

	// Retrieve database credentials from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSslMode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", dbHost, dbUser, dbPassword, dbName, dbPort, dbSslMode)

	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to the database: %v", err)
	}

	fmt.Println("Conected to the database.........")
	return d
}

var db *gorm.DB = Connect()
