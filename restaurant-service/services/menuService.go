package services

import (
	"github.com/subhankardas/go-microservices/restaurant-service/core"
	"github.com/subhankardas/go-microservices/restaurant-service/data"
	"github.com/subhankardas/go-microservices/restaurant-service/models"
)

type MenuService interface {
	GetAllMenu(trxId string) ([]models.Menu, error)
	AddMenu(trxId string, menu *models.Menu) error
}

type menuService struct {
	log  core.Logger
	data data.MenuData
}

func NewMenuService(logger core.Logger) MenuService {
	return &menuService{
		log:  logger,
		data: data.NewMenuData(logger),
	}
}

// Implementations for MenuService interface //

func (service *menuService) GetAllMenu(trxId string) ([]models.Menu, error) {
	return service.data.GetAllMenu(trxId)
}

func (service *menuService) AddMenu(trxId string, menu *models.Menu) error {
	return service.data.AddMenu(trxId, menu)
}
