package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	r "forum/Routes"
)

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
		port = "1919"
	}
	log.Println("Server running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
