package controllers

import (
	"bytes"
	"encoding/json"
	"io"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/subhankardas/go-microservices/restaurant-service/core"
	"github.com/subhankardas/go-microservices/restaurant-service/data"
	"github.com/subhankardas/go-microservices/restaurant-service/models"
	"github.com/subhankardas/go-microservices/restaurant-service/services"
)

var config *models.Config
var logger core.Logger
var router *gin.Engine
var menuService services.MenuService

// Setup required dependencies here.
func Setup() {
	// Load config properties
	config = core.LoadConfig("../configs/", "qa.config", "yml")

	// Create new logger
	logger = core.NewLogger(config.Log)

	router = gin.Default()

	// Initialize required dependencies for controller
	db := core.NewDatabase(config, logger)
	menuData := data.NewMenuData(config, logger, db)
	menuService = services.NewMenuService(config, logger, menuData)
}

// Test GET endpoint for fetching list of all menus.
func TestGetAllMenu(t *testing.T) {
	t.Run("api_returns_list_of_all_menus", func(t *testing.T) {
		Setup()

		// Given - Setup controller and route
		menuCtrl := NewMenuController(config, logger, menuService)
		router.GET("/api/menu", menuCtrl.GetAllMenu)

		// When - Make http request and record response
		request, _ := http.NewRequest(http.MethodGet, "/api/menu", nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		var data []models.Menu
		body, err1 := io.ReadAll(response.Body)
		err2 := json.Unmarshal(body, &data)

		// Then - Assert data and errors
		assert.Equal(t, http.StatusOK, response.Code)
		assert.ErrorIs(t, err1, nil)
		assert.ErrorIs(t, err2, nil)
		assert.GreaterOrEqual(t, len(data), 0)
	})
}

func TestAddMenu(t *testing.T) {
	t.Run("api_adds_new_menu_details", func(t *testing.T) {
		Setup()

		// Given - Setup controller and route
		menuCtrl := NewMenuController(config, logger, menuService)
		router.POST("/api/menu", menuCtrl.AddMenu)

		menu := models.Menu{
			Title: "Test Menu",
			Items: []models.Item{
				{
					Name:        "Test Item 1",
					Price:       1.234,
					Description: "Test Description 1",
				},
			},
		}

		data, _ := json.Marshal(menu)

		// When - Make http request and record response
		request, _ := http.NewRequest(http.MethodPost, "/api/menu", bytes.NewBuffer(data))
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		var menuData models.Menu
		body, err1 := io.ReadAll(response.Body)
		err2 := json.Unmarshal(body, &menuData)

		// Then - Assert data and errors
		assert.Equal(t, http.StatusCreated, response.Code)
		assert.ErrorIs(t, err1, nil)
		assert.ErrorIs(t, err2, nil)
		assert.Equal(t, 32, len(menuData.ID))
	})
}
