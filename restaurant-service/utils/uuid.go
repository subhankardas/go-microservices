package utils

import (
	"strings"

	"github.com/google/uuid"
)

func NewUUID() string {
	return uuid.New().String()
}

func NewID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
