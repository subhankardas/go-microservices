package services

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	cmocks "github.com/subhankardas/go-microservices/restaurant-service/mocks/core"
	mocks "github.com/subhankardas/go-microservices/restaurant-service/mocks/data"
	"github.com/subhankardas/go-microservices/restaurant-service/models"
)

const (
	mockTrxId = "1234"
)

func TestAddMenu(t *testing.T) {
	t.Run("add_menu_successfully", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Given - setup mock calls, expectations and dummy data
		mockMenu := models.Menu{Title: "Test Menu", Items: []models.Item{{Name: "Item 1", Description: "Sample."}}}
		mockMenuData := mocks.NewMockMenuData(ctrl)
		mockMenuData.EXPECT().AddMenu(mockTrxId, &mockMenu).Return(nil)

		// When - add menu function is called
		menuService := NewMenuService(&models.Config{}, cmocks.NewMockLogger(ctrl), mockMenuData)
		err := menuService.AddMenu(mockTrxId, &mockMenu)

		// Then - assert actual data and errors
		assert.Equal(t, nil, err)
		assert.NotNil(t, mockMenu)
		assert.Len(t, mockMenu.ID, 32)
	})
}
