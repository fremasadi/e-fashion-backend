package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"go-mysql-backend/models"

	"github.com/gorilla/mux"
)

func GetOrderDetails(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM order_details")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		orderDetails := []models.OrderDetail{}
		for rows.Next() {
			var orderDetail models.OrderDetail
			var createdAt []byte
			err := rows.Scan(&orderDetail.ID, &orderDetail.UserID, &orderDetail.ProductID, &orderDetail.Quantity, &orderDetail.Price, &orderDetail.Size, &orderDetail.Color, &createdAt)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// Convert []byte to time.Time
			orderDetail.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			orderDetails = append(orderDetails, orderDetail)
		}
		json.NewEncoder(w).Encode(orderDetails)
	}
}

func GetOrderDetail(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		var orderDetail models.OrderDetail
		var createdAt []byte
		err := db.QueryRow("SELECT * FROM order_details WHERE id = ?", id).Scan(&orderDetail.ID, &orderDetail.UserID, &orderDetail.ProductID, &orderDetail.Quantity, &orderDetail.Price, &orderDetail.Size, &orderDetail.Color, &createdAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Convert []byte to time.Time
		orderDetail.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(orderDetail)
	}
}

func CreateOrderDetail(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var orderDetail models.OrderDetail
		err := json.NewDecoder(r.Body).Decode(&orderDetail)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		result, err := db.Exec("INSERT INTO order_details (user_id, product_id, quantity, price, size, color) VALUES (?, ?, ?, ?, ?, ?)", orderDetail.UserID, orderDetail.ProductID, orderDetail.Quantity, orderDetail.Price, orderDetail.Size, orderDetail.Color)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		id, err := result.LastInsertId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		orderDetail.ID = int(id)
		json.NewEncoder(w).Encode(orderDetail)
	}
}

func UpdateOrderDetail(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		var updateData map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&updateData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Assuming only quantity is being updated
		quantity, ok := updateData["quantity"].(float64)
		if !ok {
			http.Error(w, "Invalid quantity", http.StatusBadRequest)
			return
		}

		_, err = db.Exec("UPDATE order_details SET quantity = ? WHERE id = ?", quantity, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent) // No content for successful update
	}
}

func DeleteOrderDetail(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		_, err := db.Exec("DELETE FROM order_details WHERE id = ?", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
