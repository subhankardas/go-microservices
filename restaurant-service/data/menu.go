package data

import "github.com/subhankardas/go-microservices/restaurant-service/models"

func GetAllMenu() models.Menu {
	return models.Menu{
		Title: "Breakfast Menu",
		Items: []*models.Item{
			{
				Name:        "Sandwich",
				Description: "Some text!",
				Price:       1.2,
			},
		},
	}
}
