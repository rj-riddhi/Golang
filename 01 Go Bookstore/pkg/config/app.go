// // config package handles database connections using GORM.
// package config

// import (
// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/mysql" // Import MySQL dialect for GORM
// )

// var (
// 	db *gorm.DB // Global variable to hold the database connection
// )

// // Connect initializes a connection to the MySQL database.
// // It uses GORM to open a connection with the given DSN (Data Source Name).
// func Connect() {
// 	// Open a connection to the database using GORM with the provided DSN
// 	// d, err := gorm.Open("mysql", "postgres:password1@tcp(localhost:3306)/bookstore?charset=utf8&parseTime=True&loc=Local")
// 	d, err := gorm.Open("mysql", "postgres:password1@tcp(localhost:3306)/bookstore?charset=utf8&parseTime=True&loc=Local")

// 	if err != nil {
// 		// Panic if there is an error opening the database connection
// 		panic(err)
// 	}
// 	// Assign the database connection to the global db variable
// 	db = d
// }

// // GetDB returns the current database connection.
// // This function provides access to the database connection
// // for other parts of the application that require it.
// func GetDB() *gorm.DB {
// 	return db
// }

// config package handles database connections using GORM.
package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB // Global variable to hold the database connection
)

// Connect initializes a connection to the MySQL database.
// It uses GORM to open a connection with the given DSN (Data Source Name).
func Connect() {
	// Open a connection to the database using GORM with the provided DSN

	dsn := "host=localhost user=postgres password=password1 dbname=bookstore port=5432 sslmode=disable"
	// Open a connection to the PostgreSQL database using GORM
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err) // Log error and exit if connection fails
	}
	// Assign the database connection to the global db variable
	db = d
}

// GetDB returns the current database connection.
// This function provides access to the database connection
// for other parts of the application that require it.
func GetDB() *gorm.DB {
	return db
}
