package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"go-mysql-backend/models"

	"github.com/gorilla/mux"
)

func GetProducts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, type, image, original_price, discounted_price, rating, review_count, is_on_sale FROM products")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var products []models.Product
		for rows.Next() {
			var product models.Product
			err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Image, &product.OriginalPrice, &product.DiscountedPrice, &product.Rating, &product.ReviewCount, &product.IsOnSale)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			products = append(products, product)
		}
		json.NewEncoder(w).Encode(products)
	}
}

func GetProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		var product models.Product
		err = db.QueryRow("SELECT id, name, type, image, original_price, discounted_price, rating, review_count, is_on_sale FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Type, &product.Image, &product.OriginalPrice, &product.DiscountedPrice, &product.Rating, &product.ReviewCount, &product.IsOnSale)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Product not found", http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		json.NewEncoder(w).Encode(product)
	}
}

func CreateProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product models.Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := db.Exec("INSERT INTO products (name, type, image, original_price, discounted_price, rating, review_count, is_on_sale) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", product.Name, product.Type, product.Image, product.OriginalPrice, product.DiscountedPrice, product.Rating, product.ReviewCount, product.IsOnSale)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		product.ID = int(id)
		json.NewEncoder(w).Encode(product)
	}
}

func UpdateProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStr)
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

		_, err = db.Exec("UPDATE products SET name = ?, type = ?, image = ?, original_price = ?, discounted_price = ?, rating = ?, review_count = ?, is_on_sale = ? WHERE id = ?", product.Name, product.Type, product.Image, product.OriginalPrice, product.DiscountedPrice, product.Rating, product.ReviewCount, product.IsOnSale, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(product)
	}
}

func DeleteProduct(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		_, err = db.Exec("DELETE FROM products WHERE id = ?", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
