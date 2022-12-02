package services

import (
	"github.com/subhankardas/go-microservices/restaurant-service/core"
	"github.com/subhankardas/go-microservices/restaurant-service/data"
	"github.com/subhankardas/go-microservices/restaurant-service/models"
)

type MenuService interface {
	GetAllMenu() models.Menu
}

type menuService struct {
	log core.Logger
}

func NewMenuService(logger core.Logger) MenuService {
	return &menuService{
		log: logger,
	}
}

// Implementations for MenuService interface //

func (service *menuService) GetAllMenu() models.Menu {
	service.log.Debug("Getting menu details from DB.")
	return data.GetAllMenu()
}
