package models

import "time"

type OrderDetail struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	ProductID int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	Size      string    `json:"size"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
}
