package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func dbInit() error {
	datastoreName := os.Getenv("POSTGRES_CONNECTION")

	var err error
	db, err = sql.Open("postgres", datastoreName)
	if err != nil {
		log.Fatal(err)
	}

	if err := createTable(); err != nil {
		log.Fatal(err)
	}
	return err
}

func createTable() error {
	stmt := `CREATE TABLE IF NOT EXISTS books (
			name	VARCHAR(255),
			comment VARCHAR(255)
		)`
	_, err := db.Exec(stmt)
	return err
}

type Book struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

func bookRegister(w http.ResponseWriter, req *http.Request) {
	var b Book

	err := json.NewDecoder(req.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Person: %+v", b)

	if err := recordBook(b.Name, b.Comment); err != nil {
		msg := fmt.Sprintf("Could not save book: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "\nSuccessfully stored an entry of the current request.")

}

func recordBook(name, comment string) error {
	stmt := "INSERT INTO books (name, comment) VALUES ($1, $2)"
	_, err := db.Exec(stmt, name, comment)
	return err
}

func bookList(w http.ResponseWriter, req *http.Request) {
	books, err := queryBooks(5)
	if err != nil {
		msg := fmt.Sprintf("Could not get books: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(books)
	if err != nil {
		msg := fmt.Sprintf("Could not marshal books: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)

}

func queryBooks(limit int) ([]Book, error) {
	rows, err := db.Query("SELECT name, comment FROM books ORDER BY name DESC LIMIT $1", limit)
	if err != nil {
		return nil, fmt.Errorf("Could not get recent books: %v", err)
	}
	defer rows.Close()

	var Books []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.Name, &b.Comment); err != nil {
			return nil, fmt.Errorf("Could not get name/comment out of row: %v", err)
		}
		Books = append(Books, b)
	}
	return Books, rows.Err()
}
