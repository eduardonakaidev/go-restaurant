package models

import "github.com/google/uuid"

type User struct {
	ID            uuid.UUID `json:"id"`
	Name          string `json:"name" validate:"required"`
	Email         string `json:"email" validate:"required"`
	Password_hash string `json:"password" validate:"required"`
	Role          string `json:"role" validate:"required"`
}

// ALTER TABLE
//     "service" ADD PRIMARY KEY("id");
// CREATE TABLE "user"(
//     "id" UUID NOT NULL,
//     "name" VARCHAR(255) NOT NULL,
//     "email" VARCHAR(255) NOT NULL,
//     "password_hash" VARCHAR(255) NOT NULL,
//     "role" VARCHAR(255) NOT NULL
// );
func ValidateUser(user *User) error {
	return validate.Struct(user)
}
