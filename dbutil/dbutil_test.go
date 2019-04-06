package dbutil

import (
	"context"
	"testing"
)

type MockCursor struct {
	Count int
	Data  []map[string]interface{}
}

func (m *MockCursor) Next(context context.Context) bool {
	return m.Count > 0
}

func (m *MockCursor) Decode(val interface{}) error {
	val = m.Data[m.Count-1]
	m.Count--
	return nil
}

func (m *MockCursor) Close(context.Context) error { return nil }

var cursor = &MockCursor{
	Count: 5,
	Data: []map[string]interface{}{
		{"name": "teste 1"},
		{"name": "teste 2"},
		{"name": "teste 3"},
		{"name": "teste 4"},
		{"name": "teste 5"},
	},
}

func TestGetCollectionWithoutError(t *testing.T) {
	_, err := getCollection("books")
	if err != nil {
		t.Error(
			"Expect:", "Not throw an error",
			"Got:", err,
		)
	}
}

func TestGetCollectionValid(t *testing.T) {
	c, _ := getCollection("books")
	if c == nil {
		t.Error(
			"Expect:", "A valid collection",
			"Got:", c,
		)
	}
}

func TestGetCollectionEmpty(t *testing.T) {
	_, err := getCollection("")
	if err == nil || err.Error() != "getCollection: name is required" {
		t.Error(
			"Expect:", "Required param error",
			"Got:", err,
		)
	}
}

func TestGetAllWithoutError(t *testing.T) {
	repo := &Repository{CollectionName: "books"}
	_, err := repo.GetAll()
	if err != nil {
		t.Error(
			"Expect:", "Not throw an error",
			"Got:", err,
		)
	}
}

func TestGetAll(t *testing.T) {
	repo := &Repository{CollectionName: "books"}
	documents, _ := repo.GetAll()
	if documents == nil {
		t.Error(
			"Expect:", "A valid array of documents",
			"Got:", documents,
		)
	}
}

func TestParseCursor(t *testing.T) {
	docs, _ := parseCursor(context.TODO(), cursor)
	if docs == nil {
		t.Error(
			"Expected:", "A valid array of parsed docs",
			"Got:", docs,
		)
	}
}

func TestParseCursorWithoutError(t *testing.T) {
	_, err := parseCursor(context.TODO(), cursor)
	if err != nil {
		t.Error(
			"Expected:", "Not throw an error",
			"Got:", err,
		)
	}
}
