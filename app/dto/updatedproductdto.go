package dto



type UpdateProductDTO struct {
	Name         string    `json:"name" validate:"required"`
	Description  string    `json:"description" validate:"required"`
	PriceInCents int64     `json:"priceincents" validate:"required"`
	Stock        int64     `json:"stock" validate:"required"`
	IsActive     bool      `json:"isActive"  validate:"required"`
}