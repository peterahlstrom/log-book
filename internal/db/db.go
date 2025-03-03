package db

import (
	"context"

	"github.com/peterahlstrom/log-book/internal/models"
)

type Database interface {
	AddBook(ctx context.Context, book models.Book) (*int, error)
	GetAllBooks(ctx context.Context) ([]models.BookSummary, error)
	GetBookById(ctx context.Context, id string) (models.Book, error)
	// DeleteBook(ctx context.Context, id string) error
}
