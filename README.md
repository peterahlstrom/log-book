# log-book

## Prerequisites

Go version > 1.23.5

## Development

### Tests

Run tests:

```bash
$ go test ./...
```

## Deployment

Edit `config.json`

### config.json example

```json
{
  "validApiKeys": {
    "abc123": "dev",
    "def456": "prod"
  }
}
```

Deploy server and database

```bash
$ docker compose up
```

## Endpoints

### POST /book

Adds a new book to the database.

- The request must contain a valid JSON payload.

#### Example Request Body:

```json
{
  "title": "Confederacy of Dunces",
  "author": "John Kennedy Toole",
  "year": 1980,
  "publisher": "Louisiana State University",
  "readtime": "2023-11-24",
  "rating": 4,
  "comments": "",
  "language": "English",
  "genre": "Satire",
  "isbn": "0-8071-0657-7"
}
```

---

### GET /book

Returns a list of all books.

#### Example Response:

```json
[
  {
    "id": "1",
    "title": "Foo Bar",
    "author": "Foo"
  },
  {
    "id": "2",
    "title": "Confederacy of Dunces",
    "author": "John Kennedy Toole"
  }
]
```

---

### GET /book/:id

Returns a book by its unique identifier.

- `:id` is the unique book ID.

#### Example Response:

```json
{
  "id": "1",
  "title": "Foo Bar",
  "author": "Foo",
  "year": 2021,
  "publisher": "Example Publisher",
  "readtime": "2023-12-10",
  "rating": 5,
  "comments": "Great read!",
  "language": "English",
  "genre": "Fiction",
  "isbn": "123-456-789"
}
```

---

### DELETE /book/:id

Deletes a book from the database.

`:id` is the unique book ID.
