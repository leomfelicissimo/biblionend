package repositories

import (
	"log"

	"github.com/leomfelicissimo/biblionend/dbutil"
)

// Book represents a basic structure of book data
type Book struct {
	Name         string
	Abbreviation string
}

// BookRepository represents a repository to manage books documents
type BookRepository struct{}

func documentToBook(document map[string]interface{}) Book {
	return Book{
		Name:         document["name"].(string),
		Abbreviation: document["abbreviation"].(string),
	}
}

func parseDocuments(documents []map[string]interface{}) []Book {
	var books []Book
	for _, document := range documents {
		book := documentToBook(document)
		books = append(books, book)
	}

	return books
}

// GetAll gets all data from books collection
func (r BookRepository) GetAll() ([]Book, error) {
	log.Println("Getting all books")

	repository := &dbutil.Repository{CollectionName: "books"}

	bookDocuments, err := repository.GetAll()
	if err != nil {
		return nil, err
	}

	return parseDocuments(bookDocuments), nil
}
