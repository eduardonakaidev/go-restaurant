package models

import (
	"time"

	"github.com/google/uuid"
)

// product representa um product no banco de dados
type Product struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name" validate:"required"`
	Description  string    `json:"description" validate:"required"`
	PriceInCents int64     `json:"priceincents" validate:"required"`
	Stock        int64     `json:"stock" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	IsActive     bool      `json:"isActive"  validate:"required"`
}

// productInput representa um dto de request para o createproductcontroller
type UpdateProductDTO struct {
	Name         string `json:"name" validate:"required"`
	Description  string `json:"description" validate:"required"`
	PriceInCents int64  `json:"priceincents" validate:"required"`
	Stock        int64  `json:"stock" validate:"required"`
	IsActive     bool   `json:"isActive"  validate:"required"`
}

//representa uma response
type ProductOutput struct {
	Product
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
