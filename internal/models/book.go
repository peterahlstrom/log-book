package models

type Book struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Year      string `json:"year"`
	Publisher string `json:"publisher"`
	ReadTime  string `json:"readtime"`
	Rating    string `json:"rating"`
	Comments  string `json:"comments"`
	Language  string `json:"language"`
	Genre     string `json:"genre"`
	ISBN      string `json:"isbn"`
}

type BookSummary struct {
	ID     string `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
}
