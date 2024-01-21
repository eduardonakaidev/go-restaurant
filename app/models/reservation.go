package models

import (
	"time"
)

type Reservation struct {
	ID                string    `json:"id"`
	CreatedAt         time.Time `json:"created_at"`
	DateOfReservation time.Time `json:"dateofreservation"  validate:"required"`
	ReservedId        string    `json:"reservedid" validate:"required"`
	TableNumber       int       `json:"tablenumber"`
	CreatedId         string    `json:"createdid" validate:"required"`
	IsFinished        bool      `json:"isfinished"  validate:"required"`
	IsActiveNow       bool      `json:"isactivenow"  validate:"required"`
}

// ALTER TABLE
//
//	"user" ADD PRIMARY KEY("id");
//
// CREATE TABLE "reservation"(
//
//	"id" UUID NOT NULL,
//	"createdAt" DATE NOT NULL,
//	"DateOfReservation" DATE NOT NULL,
//	"ReservedId" UUID NOT NULL,
//	"tableNumber" INTEGER NOT NULL,
//	"createdId" UUID NOT NULL,
//	"isFinished" BOOLEAN NOT NULL,
//	"isActiveNow" BOOLEAN NOT NULL
//
// );
func ValidateReservation(reservation *Reservation) error {
	return validate.Struct(reservation)
}
