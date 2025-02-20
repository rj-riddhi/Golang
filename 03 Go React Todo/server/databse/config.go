package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnv() (err error) {

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	return err

}

func Connect() *gorm.DB {

	LoadEnv()

	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")
	db_sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", db_host, db_user, db_password, db_name, db_port, db_sslmode)

	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("error to connecting db %v \n", err)
	}

	fmt.Println("Conneted to the database")

	return d

}

var db *gorm.DB = Connect()
