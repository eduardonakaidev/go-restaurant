package utils

import (
	"strings"

	"github.com/aidarkhanov/nanoid"
)

func NewUUID() string {
	id := nanoid.New()
	id = strings.Replace(id, "-", "", -1)
	id = strings.Replace(id, "_", "", -1)
	id = strings.ToUpper(id)
	return id
}