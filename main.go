package main

import (
	"database/sql"
	"log"
	"net/http"

	"go-mysql-backend/handlers"
	"go-mysql-backend/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

func main() {
	log.Println("Starting the application...")

	// Configure the database connection (replace with your own credentials)
	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/productdb")
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database:", err)
	}
	log.Println("Database connection successful")

	// Initialize handlers with the database connection
	handlers.Initialize(db)
	handlers.InitializeUsers(db)
	handlers.InitializeLogin(db)

	// Set up the router
	router := mux.NewRouter()

	// Register routes
	routes.RegisterRoutes(router)

	// Start the server
	log.Println("Starting the server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
