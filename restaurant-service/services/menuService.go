package services

import (
	"github.com/subhankardas/go-microservices/restaurant-service/core"
	"github.com/subhankardas/go-microservices/restaurant-service/data"
	"github.com/subhankardas/go-microservices/restaurant-service/models"
	"github.com/subhankardas/go-microservices/restaurant-service/utils"
)

type MenuService interface {
	GetAllMenu(trxId string) ([]models.Menu, error)
	AddMenu(trxId string, menu *models.Menu) error
}

type menuService struct {
	config *models.Config
	log    core.Logger
	data   data.MenuData
}

// Constructor for menu business logic service layer.
func NewMenuService(config *models.Config, logger core.Logger) MenuService {
	return &menuService{
		config: config,
		log:    logger,
		data:   data.NewMenuData(config, logger),
	}
}

// Implementations for MenuService interface //

func (service *menuService) GetAllMenu(trxId string) ([]models.Menu, error) {
	return service.data.GetAllMenu(trxId)
}

func (service *menuService) AddMenu(trxId string, menu *models.Menu) error {
	// Initialize menu ID
	menu.ID = utils.NewID()

	return service.data.AddMenu(trxId, menu)
}
