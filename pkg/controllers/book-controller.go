package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Sparsh017-byte/book-management-system/pkg/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	books, err := models.GetAllBooks()
	if err != nil {
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	id, err := getBookID(r)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, err := models.GetBookById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to fetch book", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if book.Name == "" || book.Author == "" || book.Publication == "" {
		http.Error(w, "Name, author, and publication are required", http.StatusBadRequest)
		return
	}

	createdBook, err := book.CreateBook()
	if err != nil {
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, createdBook)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := getBookID(r)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var updates models.Book
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	book, err := models.GetBookById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to fetch book", http.StatusInternalServerError)
		return
	}

	if updates.Name != "" {
		book.Name = updates.Name
	}
	if updates.Author != "" {
		book.Author = updates.Author
	}
	if updates.Publication != "" {
		book.Publication = updates.Publication
	}

	result := models.GetDB().Save(book)
	if result.Error != nil {
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := getBookID(r)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, err := models.DeleteBook(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, book)
}

func getBookID(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["bookId"], 10, 64)
	if err != nil || id <= 0 {
		return 0, errors.New("invalid book ID")
	}

	return id, nil
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {

		return
	}
}
