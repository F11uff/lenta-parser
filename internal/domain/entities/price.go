
package entities

import (
	"time"
)

type Price struct {
	ID          string    `json:"id"`
	ProductID   string    `json:"product_id"`
	Current     float64   `json:"current"`
	Original    float64   `json:"original"`
	Currency    string    `json:"currency"`
	Unit        string    `json:"unit"`
	Discount    float64   `json:"discount"`
	IsPromo     bool      `json:"is_promo"`
	PromoText   string    `json:"promo_text"`
	CardPrice   float64   `json:"card_price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}