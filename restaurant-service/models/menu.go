package models

type Item struct {
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}

type Menu struct {
	Base
	Title string  `json:"title"`
	Items []*Item `json:"items"`
}
