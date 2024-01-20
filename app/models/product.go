package models

type Product struct {
	ID           string `gorm:"type:varchar(255);primary_key" json:"id"`
	Name         string `gorm:"type:varchar(255)" json:"name" validate:"required"`
	Description  string `gorm:"type:varchar(255)" json:"description" validate:"required"`
	PriceInCents int64 `gorm:"type:integer" json:"price" validate:"required"`
}

func ValidateProduct(product *Product) error {
	return validate.Struct(product)
}