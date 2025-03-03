package mocks

import (
	"context"
	"fmt"

	"github.com/peterahlstrom/log-book/internal/db"
	"github.com/peterahlstrom/log-book/internal/models"
)

type MockDB struct {
	Book  models.Book
	Books []models.BookSummary
}

var _ db.Database = (*MockDB)(nil)

func (m *MockDB) GetAllBooks(ctx context.Context) ([]models.BookSummary, error) {
	return m.Books, nil
}

func (m *MockDB) AddBook(ctx context.Context, book models.Book) (*int, error) {
	id := 1
	return &id, nil
}

func (m *MockDB) GetBookById(ctx context.Context, id string) (models.Book, error) {
	if id != "1" {
		return models.Book{}, fmt.Errorf("not found")
	}
	return m.Book, nil
}
