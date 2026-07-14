package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

func main() {
	// Configure your PostgreSQL database details here
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// New logger for detailed SQL logging
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger, // add Logger
	})
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Print(db)
	db.AutoMigrate(&Book{})
	// db.Migrator().DropColumn(&Book{}, "name")

	// Create Book
	// newBook := Book{Name: "SEOBOOK", Author: "poohlikung", Description: "Hello world12312", price: 200}
	// createBook(db, &newBook)

	// getBook
	// currentBook := getBook(db, 2)

	// updateBook
	// currentBook.Author = "Heee"
	// currentBook.price = 4000
	// updateBook(db, currentBook)
	// fmt.Print(getBook(db, 1))

	// soft Delete
	// deleteBook(db,2)
	// hard Delete
	// hardDeleteBook(db, 2)

	// search Book by name
	// fmt.Print(searchBookName(db, "SEOBOOK"))
}
