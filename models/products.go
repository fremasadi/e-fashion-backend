package models

type Product struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	Image           string  `json:"image"`
	OriginalPrice   float64 `json:"original_price"`
	DiscountedPrice float64 `json:"discounted_price,omitempty"`
	Rating          float64 `json:"rating"`
	ReviewCount     int     `json:"review_count"`
	IsOnSale        bool    `json:"is_on_sale"`
}
