package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	env "github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/subhankardas/go-microservices/restaurant-service/core"
	"github.com/subhankardas/go-microservices/restaurant-service/models"
)

var config *models.Config
var logger core.Logger
var router *gin.Engine

// Setup required dependencies here.
func Setup() {
	// Load config properties
	config = core.LoadConfig("../config.yml")

	// Create new logger
	logger = core.NewLogger(config.Log)

	// Load environment variables
	if err := env.Load("../.env"); err != nil {
		fmt.Println("Error loading .env file")
	}

	router = gin.Default()
}

// Test GET endpoint for fetching list of all menus.
func TestGetAllMenu(t *testing.T) {
	t.Run("api_returns_list_of_all_menus", func(t *testing.T) {
		Setup()

		// GIVEN
		menuCtrl := NewMenuController(config, logger)
		router.GET("/api/menu", menuCtrl.GetAllMenu)

		// WHEN
		request, _ := http.NewRequest(http.MethodGet, "/api/menu", nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		var data []models.Menu
		body, err1 := io.ReadAll(response.Body)
		err2 := json.Unmarshal(body, &data)

		// THEN
		assert.Equal(t, http.StatusOK, response.Code)
		assert.ErrorIs(t, err1, nil)
		assert.ErrorIs(t, err2, nil)
		assert.GreaterOrEqual(t, len(data), 0)
	})
}

func TestAddMenu(t *testing.T) {
	t.Run("api_adds_new_menu_details", func(t *testing.T) {
		Setup()

		// GIVEN
		menuCtrl := NewMenuController(config, logger)
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

		// WHEN
		request, _ := http.NewRequest(http.MethodPost, "/api/menu", bytes.NewBuffer(data))
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		var menuData models.Menu
		body, err1 := io.ReadAll(response.Body)
		err2 := json.Unmarshal(body, &menuData)

		// THEN
		assert.Equal(t, http.StatusCreated, response.Code)
		assert.ErrorIs(t, err1, nil)
		assert.ErrorIs(t, err2, nil)
		assert.Equal(t, 32, len(menuData.ID))
	})
}
