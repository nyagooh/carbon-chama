package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	r "forum/Routes"
)

func InitializeDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	
		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY NOT NULL,
			username TEXT UNIQUE,
			password TEXT,
			email TEXT UNIQUE
		);
		CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);
		CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
	`)
	if err != nil {
		return nil, fmt.Errorf("error creating users table: %v", err)
	}
	return db, nil
}

func main() {
	if len(os.Args) != 1 {
		fmt.Println("invalid number of arguments.")
		fmt.Println("Usage: go run .")
		return
	}

	db, err := InitializeDB()
	if err != nil {
		fmt.Printf("failed to initialize database: %v", err)
	}
	defer db.Close()

	r := r.Router()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server running on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
