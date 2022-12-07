package utils

import (
	"strings"

	"github.com/google/uuid"
)

// Create new UUID, mainly used as transaction ID.
func NewUUID() string {
	return uuid.New().String()
}

// Create new ID, mainly used as primary key ID.
func NewID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
