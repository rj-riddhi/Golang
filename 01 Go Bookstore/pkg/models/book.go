// models package defines the data model for the application and provides methods to interact with the database.
package models

import (
	"github.com/radhika.parmar/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB // Global variable to hold the database connection

// Book represents the structure of a book in the database.
type Book struct {
	gorm.Model
	Name        string `json:"name"`        // Name of the book
	Author      string `json:"author"`      // Author of the book
	Publication string `json:"publication"` // Publication of the book
}

// init initializes the database connection and migrates the Book model to the database.
// This function is automatically called when the package is imported.
func init() {
	config.Connect()        // Establish a connection to the database
	db = config.GetDB()     // Retrieve the database connection
	db.AutoMigrate(&Book{}) // Automatically migrate the Book model to the database
}

// CreateBook creates a new record for the Book in the database.
// @return *Book - the created book with its ID populated.
func (b *Book) CreateBook() *Book {
	db.Create(&b) // Create the book record in the database
	return b      // Return the created book
}

// GetAllBook retrieves all books from the database.
// @return []Book - a slice of all books.
func GetAllBook() []Book {
	var Books []Book
	db.Find(&Books) // Fetch all books from the database
	return Books    // Return the list of books
}

// GetBookById retrieves a book by its ID from the database.
// @param Id int64 - the ID of the book to retrieve.
// @return (*Book, *gorm.DB) - the retrieved book and the database instance.
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", Id).Find(&getBook) // Find the book by ID
	return &getBook, db                         // Return the book and the database instance
}

// DeleteBook deletes a book by its ID from the database.
// @param ID int64 - the ID of the book to delete.
// @return Book - the deleted book (with zero values as it's removed from DB).
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID = ?", ID).Delete(&book) // Delete the book with the specified ID
	return book                          // Return the deleted book
}
