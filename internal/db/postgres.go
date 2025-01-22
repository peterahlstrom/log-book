package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/peterahlstrom/log-book/internal/models"
)

type PostgresDB struct {
	Conn *sql.DB
}

var _ Database = (*PostgresDB)(nil)

func (p *PostgresDB) GetAllBooks(ctx context.Context) ([]models.BookSummary, error) {
	query := "SELECT id, author, title FROM logbook.books"

	rows, err := p.Conn.Query(query)
	if err != nil {
		log.Printf("ERROR: could not query all books. %v", err)
		return nil, fmt.Errorf("Database query error")
	}
	defer rows.Close()

	var books []models.BookSummary
	for rows.Next() {
		var b models.BookSummary
		err := rows.Scan(&b.ID, &b.Author, &b.Title)
		if err != nil {
			log.Printf("ERROR: failed to fetch db row. %v", err)
			return nil, fmt.Errorf("Database error")
		}
		books = append(books, b)
	}
	return books, nil
}

func (p *PostgresDB) AddBook(ctx context.Context, b models.Book) (*int, error) {

	query := `
			INSERT INTO logbook.books(title, author, year, publisher, readtime, rating, comments, language, genre, isbn)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
			RETURNING id;
	`

	var id int
	err := p.Conn.QueryRow(query, b.Title, b.Author, b.Year, b.Publisher, b.ReadTime, b.Rating, b.Comments, b.Language, b.Genre, b.ISBN).Scan(&id)
	if err != nil {
		log.Printf("could not write book to database: %v", err)
		return nil, fmt.Errorf("Database insert error")
	}
	return &id, nil
}

func (p *PostgresDB) GetBookById(ctx context.Context, id string) (models.Book, error) {
	var out models.Book

	query := "SELECT id, title, author, year, publisher, readtime, rating, comments, language, genre, isbn FROM logbook.books WHERE id=$1;"

	err := p.Conn.QueryRow(query, id).Scan(&out.ID, &out.Title,
		&out.Author, &out.Year, &out.Publisher, &out.ReadTime,
		&out.Rating, &out.Comments, &out.Language, &out.Genre, &out.ISBN)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Book{}, fmt.Errorf("book not found: %v", err)
		}
		return models.Book{}, fmt.Errorf("general error %v", err)
	}

	return out, nil
}
