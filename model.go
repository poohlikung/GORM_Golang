package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string
	Author      string
	Description string
	Price       uint
}

func createBook(db *gorm.DB, book *Book) {
	result := db.Create(book)

	if result.Error != nil {
		log.Fatalf("Error creating book :%v", result.Error)
	}
	fmt.Println("Create Book Success")
}

func getBook(db *gorm.DB, id uint) *Book {
	var book Book
	result := db.First(&book, id)

	if result.Error != nil {
		log.Fatalf("Error get Book : %v", result.Error)
	}
	return &book
}

func updateBook(db *gorm.DB, book *Book) {
	result := db.Save(&book)
	if result.Error != nil {
		log.Fatalf("Error update fail : %v", result.Error)
	}
	fmt.Print("Update Book success")
}

// Soft Delete
func deleteBook(db *gorm.DB, id uint) {
	var book Book
	result := db.Delete(&book, id)
	if result.Error != nil {
		log.Fatalf("Delete Book fail : %v", result.Error)
	}
	fmt.Print("Delete Book success")
}

// Hard Delete
func hardDeleteBook(db *gorm.DB, id uint) {
	var book Book
	result := db.Unscoped().Delete(&book, id)
	if result.Error != nil {
		log.Fatalf("Delete Book fail : %v", result.Error)
	}
	fmt.Print("Delete Book success")
}

// search name
func searchBookName(db *gorm.DB, bookName string) ([]Book, error) {
	var books []Book
	result := db.Where("name = ?", bookName).Find(&books)
	if result.Error != nil {
		return nil, fmt.Errorf("search book fail: %w", result.Error)
	}
	return books, nil
}
