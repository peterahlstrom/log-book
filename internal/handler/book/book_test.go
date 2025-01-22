package book_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/peterahlstrom/log-book/internal/handler/book"
	"github.com/peterahlstrom/log-book/internal/mocks"
	"github.com/peterahlstrom/log-book/internal/models"
)

var router *http.ServeMux

func TestMain(m *testing.M) {
	router = http.NewServeMux()

	router.HandleFunc("POST /book", bookService.AddBookHandler)
	router.HandleFunc("GET /book", bookService.GetAllBooksHandler)
	router.HandleFunc("GET /book/{id}", bookService.GetBookByIdHandler)
	// router.HandleFunc("DELETE /book/{id}", bookService.DeleteBookHandler)

	m.Run()
}

var mockDB = &mocks.MockDB{
	Books: []models.BookSummary{
		{ID: "1", Title: "Foo Bar", Author: "Foo"},
	},
	Book: models.Book{
		ID: "1", Title: "Testbook by id", Author: "Foo",
	},
}
var bookService = book.NewBookService(mockDB)

func TestGetAllBooks(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/book", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("got status %d, want status %d", rr.Code, http.StatusOK)
	}
}

func TestAddBook(t *testing.T) {
	reqBody := `{
    "title": "Confederation of Dunces",
    "author": "John Kennedy Tool",
    "year": "1980",
    "publisher": "Louisiana State Uni",
    "readtime": "2023-11-24",
    "rating": "4",
    "comments": "",
    "language": "English",
    "genre": "",
    "isbn": "0-8071-0657-7"
}`
	req := httptest.NewRequest(http.MethodPost, "/book", strings.NewReader(reqBody))
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("got status %d, want status %d", rr.Code, http.StatusOK)
	}
}

func TestGetBookById(t *testing.T) {
	tests := []struct {
		name       string
		id         int
		wantStatus int
	}{
		{"valid id", 1, 200},
		{"invalid id", 42, 404},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := fmt.Sprintf("/book/%d", tt.id)
			req := httptest.NewRequest(http.MethodGet, path, nil)
			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)
			if rr.Code != tt.wantStatus {
				t.Errorf("got status %d, want status %d", rr.Code, tt.wantStatus)
			}

		})
	}
}
