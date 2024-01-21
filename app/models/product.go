package models

import "time"

type Product struct {
	ID           string    `json:"id"`
	Name         string    `json:"name" validate:"required"`
	Description  string    `json:"description" validate:"required"`
	PriceInCents int64     `json:"priceincents" validate:"required"`
	Stock        int64     `json:"stock" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	IsActive     bool      `json:"isActive"  validate:"required"`
}

// "id" UUID NOT NULL,
// "name" VARCHAR(255) NOT NULL,
// "description" VARCHAR(255) NOT NULL,
// "priceInCents" INTEGER NOT NULL,
// "stock" INTEGER NOT NULL,
// "createdAt" DATE NOT NULL,
// "updatedAt" DATE NULL,
// "isActive" BOOLEAN NOT NULL
func ValidateProduct(product *Product) error {
	return validate.Struct(product)
}
