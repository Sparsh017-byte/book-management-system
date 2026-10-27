package models

import (
	"github.com/Sparsh017-byte/book-management-system/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func Init() error {
	config.Connect()
	db = config.GetDB()

	return db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() (*Book, error) {
	result := db.Create(b)
	if result.Error != nil {
		return nil, result.Error
	}

	return b, nil
}

func GetAllBooks() ([]Book, error) {
	var books []Book

	result := db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func GetBookById(id int64) (*Book, error) {
	var book Book

	result := db.First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}

func DeleteBook(id int64) (*Book, error) {
	book, err := GetBookById(id)
	if err != nil {
		return nil, err
	}

	result := db.Delete(book)
	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}

func GetDB() *gorm.DB {
	return db
}
