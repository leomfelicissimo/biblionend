package repositories

import "testing"

func TestGetAll(t *testing.T) {
	bookRepo := &BookRepository{}
	books, err := bookRepo.GetAll()

	if err != nil {
		t.Error("Error calling GetAllBooks", err)
	}

	if len(books) == 0 {
		t.Error("Books is empty")
	}
}
