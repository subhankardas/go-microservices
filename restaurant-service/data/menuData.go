package data

import (
	"github.com/subhankardas/go-microservices/restaurant-service/core"
	"github.com/subhankardas/go-microservices/restaurant-service/models"
)

type MenuData interface {
	GetAllMenu(trxId string) ([]models.Menu, error)
	AddMenu(trxId string, menu *models.Menu) error
	GetMenu(trxId string, menuId string, menu *models.Menu) error
	UpdateMenu(trxId string, menu *models.Menu) error
	DeleteMenu(trxId string, menuId string) error
}

type menuData struct {
	config *models.Config
	log    core.Logger
	db     core.Database
}

// Constructor for menu data layer.
func NewMenuData(config *models.Config, logger core.Logger, db core.Database) MenuData {
	migrate(db)

	return &menuData{
		config: config,
		log:    logger,
		db:     db,
	}
}

// Migrate required models.
func migrate(db core.Database) {
	db.AutoMigrate(&models.Menu{}, &models.Item{})
}

// Implementations for MenuData interface //

func (data *menuData) GetAllMenu(trxId string) ([]models.Menu, error) {
	menus := []models.Menu{}

	// Preload items, then load all menu from DB
	if _, err := data.db.Preload("Items").Find(&menus); err != nil {
		data.log.Errorf(trxId, "error: %v, cause: %v", core.UNABLE_TO_READ_ALL_MENU_FROM_DB, err)
		return nil, core.ErrUnableToReadAllMenuFromDb
	}

	return menus, nil
}

func (data *menuData) AddMenu(trxId string, menu *models.Menu) error {
	// Add new menu details to DB
	if _, err := data.db.Create(menu); err != nil {
		data.log.Errorf(trxId, "error: %v, cause: %v", core.UNABLE_TO_ADD_MENU_TO_DB, err)
		return core.ErrUnableToAddMenuToDb
	}

	return nil
}

func (data *menuData) GetMenu(trxId string, menuId string, menu *models.Menu) error {
	// Get first item with ID value equals menu ID
	if err := data.db.Preload("Items").First(&menu, "id = ?", menuId); err != nil {
		data.log.Errorf(trxId, "error: %v, cause: %v", core.UNABLE_TO_GET_MENU_FROM_DB, err)
		return core.ErrUnableToGetMenuFromDb
	}

	return nil
}

func (data *menuData) UpdateMenu(trxId string, menu *models.Menu) error {
	// Update existing menu details to DB
	if _, err := data.db.Save(menu); err != nil {
		data.log.Errorf(trxId, "error: %v, cause: %v", core.UNABLE_TO_UPDATE_MENU_TO_DB, err)
		return core.ErrUnableToUpdateMenuToDb
	}

	return nil
}

func (data *menuData) DeleteMenu(trxId string, menuId string) error {
	// Delete existing menu by ID
	if err := data.db.Delete(&models.Menu{}, "id = ?", menuId); err != nil {
		data.log.Errorf(trxId, "error: %v, cause: %v", core.UNABLE_TO_DELETE_MENU_FROM_DB, err)
		return core.ErrUnableToDeleteMenuFromDb
	}

	return nil
}
