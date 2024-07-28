package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"go-mysql-backend/models"

	"github.com/gorilla/mux"
)

var db *sql.DB

func Initialize(database *sql.DB) {
	db = database
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	log.Println("Fetching all products...")
	rows, err := db.Query("SELECT id, name, type, image, price, original_price, discounted_price, rating, review_count, is_on_sale FROM products")
	if err != nil {
		log.Println("Error executing query:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Image, &product.Price, &product.OriginalPrice, &product.DiscountedPrice, &product.Rating, &product.ReviewCount, &product.IsOnSale)
		if err != nil {
			log.Println("Error scanning row:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
	log.Println("Fetched all products successfully")
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var product models.Product
	err = db.QueryRow("SELECT id, name, type, image, price, original_price, discounted_price, rating, review_count, is_on_sale FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Type, &product.Image, &product.Price, &product.OriginalPrice, &product.DiscountedPrice, &product.Rating, &product.ReviewCount, &product.IsOnSale)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Product not found", http.StatusNotFound)
		} else {
			log.Println("Error executing query:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
	log.Println("Fetched product successfully")
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("Creating new product:", product.Name)
	result, err := db.Exec("INSERT INTO products (name, type, image, price, original_price, discounted_price, rating, review_count, is_on_sale) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", product.Name, product.Type, product.Image, product.Price, product.OriginalPrice, product.DiscountedPrice, product.Rating, product.ReviewCount, product.IsOnSale)
	if err != nil {
		log.Println("Error executing query:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert ID:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	product.ID = int(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
	log.Println("Created new product successfully")
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var product models.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Updating product with ID: %d\n", id)
	_, err = db.Exec("UPDATE products SET name = ?, type = ?, image = ?, price = ?, original_price = ?, discounted_price = ?, rating = ?, review_count = ?, is_on_sale = ? WHERE id = ?", product.Name, product.Type, product.Image, product.Price, product.OriginalPrice, product.DiscountedPrice, product.Rating, product.ReviewCount, product.IsOnSale, id)
	if err != nil {
		log.Println("Error executing query:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
	log.Println("Updated product successfully")
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	log.Printf("Deleting product with ID: %d\n", id)
	_, err = db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		log.Println("Error executing query:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	log.Println("Deleted product successfully")
}
