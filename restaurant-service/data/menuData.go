package data

import (
	"errors"

	"github.com/subhankardas/go-microservices/restaurant-service/core"
	"github.com/subhankardas/go-microservices/restaurant-service/models"
	"github.com/subhankardas/go-microservices/restaurant-service/utils"
)

type MenuData interface {
	GetAllMenu(trxId string) ([]models.Menu, error)
	AddMenu(trxId string, menu *models.Menu) error
}

type menuData struct {
	log core.Logger
	db  core.Database
}

func NewMenuData(logger core.Logger) MenuData {
	db := core.NewDatabase(logger)
	migrate(db)

	return &menuData{
		log: logger,
		db:  db,
	}
}

func migrate(db core.Database) {
	db.AutoMigrate(&models.Menu{}, &models.Item{})
}

// Implementations for MenuData interface //

func (data *menuData) GetAllMenu(trxId string) ([]models.Menu, error) {
	menus := []models.Menu{}
	if _, err := data.db.Preload("Items").Find(&menus); err != nil {
		data.log.Errorf(trxId, "Unable to read list of all menu from DB, error: %v", err)
		return nil, errors.New("unable to read list of all menu from DB")
	}
	return menus, nil
}

func (data *menuData) AddMenu(trxId string, menu *models.Menu) error {
	menu.ID = utils.NewID()
	if _, err := data.db.Create(&menu); err != nil {
		data.log.Errorf(trxId, "Unable to add menu details to DB, error: %v", err)
		return errors.New("unable to add menu details to DB")
	}
	return nil
}
