package models

import "time"

type Service struct {
	ID           string    ` json:"id"`
	Name         string    `json:"name" validate:"required"`
	Description  string    `json:"description" `
	PriceInCents int64     `json:"priceincents" validate:"required"`
	HasDate      bool      `json:"hasdate"`
	Date         time.Time `json:"date"`
	IsActive     bool      `json:"isActive"  validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

// ALTER TABLE
//
//	"product" ADD PRIMARY KEY("id");
//
// CREATE TABLE "service"(
//
//	"id" UUID NOT NULL,
//	"name" VARCHAR(255) NOT NULL,
//	"description" VARCHAR(255) NOT NULL,
//	"priceInCents" BIGINT NOT NULL,
//	"hasDate" BOOLEAN NOT NULL,
//	"date" DATE NOT NULL,
//	"isActive" BOOLEAN NOT NULL,
//	"createdAt" DATE NOT NULL,
//	"updatedAt" DATE NULL
//
// );
func ValidateService(service *Service) error {
	return validate.Struct(service)
}
