package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"go-mysql-backend/models"
)

var loginDb *sql.DB

func InitializeLogin(db *sql.DB) {
	loginDb = db
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials models.User
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("Attempting to log in with email:", credentials.Email)

	var user models.User
	err = loginDb.QueryRow("SELECT id, name, email, password FROM users WHERE email = ? AND password = ?", credentials.Email, credentials.Password).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Invalid credentials for email:", credentials.Email)
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		} else {
			log.Println("Error executing query:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	log.Println("Login successful for user:", user.Name)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
