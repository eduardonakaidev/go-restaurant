package models

import (
	"time"

	"github.com/google/uuid"
)

type Schema struct {
    ID        uuid.UUID        `json:"id"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
}