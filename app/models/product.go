package models

import "time"

type Product struct {
	ID           string    `gorm:"type:varchar(255);primary_key" json:"id"`
	Name         string    `gorm:"type:varchar(255)" json:"name" validate:"required"`
	Description  string    `gorm:"type:varchar(255)" json:"description" validate:"required"`
	PriceInCents int64     `gorm:"type:integer" json:"priceincents" validate:"required"`
	Stock        int64     `gorm:"type:integer" json:"stock" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	IsActive     bool      `json:"isActive"  validate:"required"`
}

func ValidateProduct(product *Product) error {
	return validate.Struct(product)
}
