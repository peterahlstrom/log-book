package book

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/peterahlstrom/log-book/internal/db"
	"github.com/peterahlstrom/log-book/internal/models"
)

type BookService struct {
	DB db.Database
}

func NewBookService(database db.Database) *BookService {
	return &BookService{DB: database}
}

func (bs *BookService) AddBookHandler(w http.ResponseWriter, r *http.Request) {
	var b *models.Book
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		log.Printf("Failed to parse request: %v", err)
		http.Error(w, "Database request error", http.StatusBadRequest)
	}
	id, err := bs.DB.AddBook(r.Context(), *b)
	if err != nil {
		fmt.Printf("ERROR: Database error: %v", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"id": %d}`, *id)))
}

func (bs *BookService) GetBookByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	book, err := bs.DB.GetBookById(r.Context(), id)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
	}

	resp, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resp))
}

func (bs *BookService) GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := bs.DB.GetAllBooks(r.Context())

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Book not found", http.StatusNotFound)
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// func (bs *BookService) DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
// 	id := r.PathValue("id")

// 	query := "DELETE FROM logbook.books WHERE id=$1;"
// 	result, err := bs.DB.Conn.Exec(query, id)
// 	if err != nil {
// 		log.Printf("ERROR: could not delete id %s: %v", id, err)
// 		http.Error(w, fmt.Sprintf("Could not delete book with id %s", id), http.StatusInternalServerError)
// 		return
// 	}
// 	count, _ := result.RowsAffected()
// 	if count > 1 {
// 		w.WriteHeader(http.StatusNotFound)
// 	}
// 	log.Printf("Deleted book with id: %s", id)
// 	w.WriteHeader(http.StatusNoContent)
// }
