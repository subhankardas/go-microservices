package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subhankardas/go-microservices/restaurant-service/utils"
)

func TestNewUUID(t *testing.T) {
	t.Run("check_new_uuid_is_valid", func(t *testing.T) {
		// When
		uuid := utils.NewUUID()

		// Then
		assert.Len(t, uuid, 36)
		assert.Contains(t, uuid, "-")
	})
}

func TestNewID(t *testing.T) {
	t.Run("check_new_id_is_valid", func(t *testing.T) {
		// When
		id := utils.NewID()

		// Then
		assert.Len(t, id, 32)
		assert.NotContains(t, id, "-")
	})
}
