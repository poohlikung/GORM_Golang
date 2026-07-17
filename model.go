package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
}

func createBook(db *gorm.DB, book *Book) error {
	result := db.Create(book)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func getBook(db *gorm.DB, id int) *Book {
	var book Book
	result := db.First(&book, id)

	if result.Error != nil {
		log.Fatalf("Error get Book : %v", result.Error)
	}
	return &book
}

// Get all book
func getBooks(db *gorm.DB) []Book {
	var books []Book
	result := db.Find(&books)

	if result.Error != nil {
		log.Fatalf("Error get Book : %v", result.Error)
	}
	return books
}

func updateBook(db *gorm.DB, book *Book) error {
	result := db.Save(&book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Soft Delete
func deleteBook(db *gorm.DB, id int) error {
	var book Book
	result := db.Delete(&book, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Hard Delete
func hardDeleteBook(db *gorm.DB, id int) error {
	var book Book
	result := db.Unscoped().Delete(&book, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// search name
func searchBookName(db *gorm.DB, bookName string) ([]Book, error) {
	var books []Book
	result := db.Where("name = ?", bookName).Order("price").Find(&books)
	if result.Error != nil {
		return nil, fmt.Errorf("search book fail: %w", result.Error)
	}
	return books, nil
}
