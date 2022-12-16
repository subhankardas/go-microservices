package models

type Menu struct {
	BaseUID
	BaseDateTimeMeta
	Title string `json:"title"`
	Items []Item `json:"items" gorm:"foreignkey:MenuID"` // One to many mapping (Menu --> []Item)
}

type Item struct {
	BaseID
	BaseDateTimeMeta
	MenuID      string  `json:"-"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}
