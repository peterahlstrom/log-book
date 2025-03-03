package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/peterahlstrom/log-book/internal/db"
	"github.com/peterahlstrom/log-book/internal/handler/book"
	"github.com/peterahlstrom/log-book/internal/utils/auth"
	"github.com/peterahlstrom/log-book/internal/utils/config"
)

var configFilePath = "config.json"

func waitForDb(conn *sql.DB, timeout int) error {
	deadline := time.After(time.Duration(timeout) * time.Second)
	pingInterval := time.Duration(1) * time.Second

	for {
		if err := conn.Ping(); err == nil {
			return nil
		}
		select {
		case <-deadline:
			return fmt.Errorf("timeout")
		case <-time.After(pingInterval):
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: ./main <port>")
	}
	port := os.Args[1]

	config, err := config.GetConfig(configFilePath)
	if err != nil {
		log.Fatalf("Could not read config file. %v", err)
	}

	dsn := "host=db user=username password=S3cret dbname=logbook sslmode=disable"

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("ERROR: failed to connect to database: %v", err)
	}
	defer conn.Close()

	if err := waitForDb(conn, 10); err != nil {
		log.Fatalf("ERROR: Timeout. Database is not responding to ping: %v", err)
	}

	postgresDB := &db.PostgresDB{Conn: conn}

	bookService := book.NewBookService(postgresDB)

	router := http.NewServeMux()

	router.HandleFunc("POST /book", bookService.AddBookHandler)
	router.HandleFunc("GET /book", bookService.GetAllBooksHandler)
	router.HandleFunc("GET /book/{id}", bookService.GetBookByIdHandler)
	// router.HandleFunc("DELETE /book/{id}", bookService.DeleteBookHandler)

	secureHandler := auth.ApiKeyMiddleware(config.ValidApiKeys)(router)

	addr := fmt.Sprintf(":%s", port)
	server := http.Server{
		Addr:    addr,
		Handler: secureHandler,
	}

	log.Printf("Starting server on port %s...\n", port)
	server.ListenAndServe()
}
