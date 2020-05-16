package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var db *sql.DB

func dbInit() error {
	datastoreName := os.Getenv("POSTGRES_CONNECTION")

	var err error
	db, err = sql.Open("posgres", datastoreName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := createTable(); err != nil {
		log.Fatal(err)
	}
	return err
}

func createTable() error {
	stmt := `CREATE TABLE IF NOT EXISTS books (
			name	VARCHAR(255)
			comment VARCHAR(255)
		)`
	_, err := db.Exec(stmt)
	return err
}

type book struct {
	name    string
	comment string
}

func bookRegister(w http.ResponseWriter, req *http.Request) {
	var b book

	err := json.NewDecoder(req.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Person: %+v", b)

	if err := recordBook(b.name, b.comment); err != nil {
		msg := fmt.Sprintf("Could not save visit: %v", err)
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
