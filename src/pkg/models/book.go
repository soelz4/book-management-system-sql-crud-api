package models

import (
	"gorm.io/gorm"

	"go-bookstore/src/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `        json:"author"`
	Publication string `        json:"publication"`
}

// The init Function is a Function that Takes no Argument and Returns Nothing.
// This Function Executes After the Package is Imported and Maintains the Order of Execution.
func init() {
	config.Connect()
	db = config.GetDb()
	// Automatically Migrate your Schema, to Keep your Schema Up to Date.
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	Books := []Book{}
	// Finds All Records that Matching with Given Condition and then ~> Put into Books Slice
	db.Find(&Books)
	return Books
}

func GetBookByID(id int64) (*Book, *gorm.DB) {
	var getBook Book
	// Finds All Records that Matching with Given Condition and then ~> Put into getBook Variable
	db := db.Where("ID=?", id).Find(&getBook)
	return &getBook, db
}

func DeleteBookByID(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(&book)
	return book
}
