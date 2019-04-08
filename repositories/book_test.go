package repositories

import "testing"

func TestGetAll(t *testing.T) {
	bookRepo := &BookRepository{}
	books, err := bookRepo.FindAll()

	if err != nil {
		t.Error("Error calling FindAll", err)
	}

	if len(books) == 0 {
		t.Error("Books is empty")
	}
}
