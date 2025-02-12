package entities

import (
	"time"
)

type ProductResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Available int       `json:"available"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

type FieldUpdatedProduct struct {
	ID        string
	Available int
}
