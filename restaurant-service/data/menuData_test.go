package data

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mocks "github.com/subhankardas/go-microservices/restaurant-service/mocks/core"
	"github.com/subhankardas/go-microservices/restaurant-service/models"
)

const (
	mockTrxId = "1234"
)

func TestGetAllMenu(t *testing.T) {
	t.Run("get_all_menu_successfully", func(t *testing.T) {
		// Setup mock controller
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Given - setup mock calls, expectations and dummy data
		mockMenu := []models.Menu{}
		mockDb := mocks.NewMockDatabase(ctrl)
		mockDb.EXPECT().AutoMigrate(&models.Menu{}, &models.Item{}).Times(1)
		mockDb.EXPECT().Preload("Items").Return(mockDb)
		mockDb.EXPECT().Find(&mockMenu).Return(int64(1), nil)

		// When - get all menu function is called
		menuData := NewMenuData(&models.Config{}, mocks.NewMockLogger(ctrl), mockDb)
		data, err := menuData.GetAllMenu(mockTrxId)

		// Then - assert actual data and errors
		assert.Equal(t, err, nil)
		assert.Equal(t, []models.Menu{}, data)
	})
}

func TestAddMenu(t *testing.T) {
	t.Run("add_menu_successfully", func(t *testing.T) {
		// Setup mock controller
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// Given - setup mock calls, expectations and dummy data
		mockMenu := models.Menu{Title: "Test Menu", Items: []models.Item{{Name: "Item 1", Description: "Sample."}}}
		mockDb := mocks.NewMockDatabase(ctrl)
		mockDb.EXPECT().AutoMigrate(&models.Menu{}, &models.Item{}).Times(1)
		mockDb.EXPECT().Create(&mockMenu).Return(int64(1), nil)

		// When - add menu function is called
		menuData := NewMenuData(&models.Config{}, mocks.NewMockLogger(ctrl), mockDb)
		err := menuData.AddMenu(mockTrxId, &mockMenu)

		// Then - assert actual data and errors
		assert.Equal(t, err, nil)
		assert.NotNil(t, mockMenu)
	})
}
