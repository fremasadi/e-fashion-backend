package main

import (
	"database/sql"
	"log"
	"net/http"

	"go-mysql-backend/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var DB *sql.DB

func main() {
	log.Println("Starting the application...")

	// Configure the database connection (replace with your own credentials)
	var err error
	DB, err = sql.Open("mysql", "root:password123@tcp(127.0.0.1:3306)/e_fashion")
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error pinging the database:", err)
	}
	log.Println("Database connection successful")

	// Set up the router
	router := mux.NewRouter()

	// Register routes
	routes.RegisterRoutes(router, DB)

	// Start the server
	log.Println("Starting the server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
