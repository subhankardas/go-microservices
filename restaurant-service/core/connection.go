package core

import (
	"fmt"

	"github.com/subhankardas/go-microservices/restaurant-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Create new PostgresDB connection.
func newPostgresDbConnection(config *models.Config) (*DB, error) {
	db := config.Database
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Africa/Lagos",
		db.Host,
		db.Username,
		db.Password,
		db.Name,
		db.Port,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
