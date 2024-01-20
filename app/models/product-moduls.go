package models


type Product struct {
    Name        string    `json:"name"`
    Description string    `json:"description"`
    PriceInCents int       `json:"priceInCents"`
    Stock       int       `json:"stock"`
	Schema
}
// entity completed
// Name        string    `json:"name"`
// Description string    `json:"description"`
// PriceInCents int       `json:"priceInCents"`
// Stock       int       `json:"stock"`
// ID        uuid.UUID        `json:"id"`
// CreatedAt time.Time      `json:"created_at"`
// UpdatedAt time.Time      `json:"updated_at"`
