package models

import "time"

type Client struct {
	ID            string    `json:"id"`
	Name          string    `json:"name" validate:"required"`
	Phone         int64     `json:"fone" validate:"required"`
	CreditInCents int64     `json:"creditincents" validate:"required"`
	Cpf           int64     `json:"cpf" validate:"required"`
	CreatedAt     time.Time `json:"created_at"`
	IsActive      bool      `json:"isActive"  validate:"required"`
}

func ValidateClient(client *Client) error {
	return validate.Struct(client)
}
