// controllers package handles HTTP requests for book-related operations.
package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/radhika.parmar/go-bookstore/pkg/models"
	"github.com/radhika.parmar/go-bookstore/pkg/utils"
)

// NewBook is a global variable that can be used to create new book instances.
var NewBook models.Book

// GetBook retrieves all books from the database and returns them in the response.
// @param w http.ResponseWriter - the response writer.
// @param r *http.Request - the incoming HTTP request.
func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBook()                    // Fetch all books
	res, _ := json.Marshal(newBooks)                   // Convert book list to JSON
	w.Header().Set("Content-Type", "application/json") // Set response header to JSON
	w.WriteHeader(http.StatusOK)                       // Set HTTP status to 200 OK
	w.Write(res)                                       // Write JSON response
}

// GetBookById retrieves a book by its ID and returns it in the response.
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                       // Get route variables
	bookId := vars["bookId"]                  // Retrieve book ID from URL
	ID, err := strconv.ParseInt(bookId, 0, 0) // Convert book ID from string to integer
	if err != nil {
		fmt.Println("error while parsing")                      // Log error if ID parsing fails
		http.Error(w, "Invalid book ID", http.StatusBadRequest) // Return 400 Bad Request if error occurs
		return
	}
	bookDetails, _ := models.GetBookById(ID) // Get book details by ID (ignoring second return value for simplicity)
	res, _ := json.Marshal(bookDetails)      // Convert book details to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// CreateBook creates a new book from the request body and saves it to the database.
func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}   // Create a new Book instance
	utils.ParseBody(r, CreateBook) // Parse the request body into the Book instance
	b := CreateBook.CreateBook()   // Save the new book to the database
	res, _ := json.Marshal(b)      // Convert saved book to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// DeleteBook deletes a book by its ID and returns a confirmation response.
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                       // Get route variables
	bookId := vars["bookId"]                  // Retrieve book ID from URL
	Id, err := strconv.ParseInt(bookId, 0, 0) // Convert book ID from string to integer
	if err != nil {
		fmt.Println("error while parsing")                      // Log error if ID parsing fails
		http.Error(w, "Invalid book ID", http.StatusBadRequest) // Return 400 Bad Request if error occurs
		return
	}
	b := models.DeleteBook(Id) // Delete the book from the database
	res, _ := json.Marshal(b)  // Convert confirmation response to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// UpdateBook updates an existing book with new details from the request body and saves changes to the database.
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}           // Create a new Book instance for updates
	utils.ParseBody(r, updateBook)            // Parse the request body into the Book instance
	vars := mux.Vars(r)                       // Get route variables
	bookId := vars["bookId"]                  // Retrieve book ID from URL
	Id, err := strconv.ParseInt(bookId, 0, 0) // Convert book ID from string to integer
	if err != nil {
		fmt.Println("error while parsing")                      // Log error if ID parsing fails
		http.Error(w, "Invalid book ID", http.StatusBadRequest) // Return 400 Bad Request if error occurs
		return
	}
	bookDetails, db := models.GetBookById(Id) // Get existing book details and database instance
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name // Update book name if provided
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author // Update book author if provided
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication // Update book publication if provided
	}
	db.Save(&bookDetails)               // Save updated book details to the database
	res, _ := json.Marshal(bookDetails) // Convert updated book details to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
