package core

import (
	"sync"

	"github.com/subhankardas/go-microservices/restaurant-service/models"
	"gorm.io/gorm"
)

type DB = gorm.DB

type Database interface {
	AutoMigrate(models ...interface{})
	Create(value interface{}) (int64, error)
	Find(destination interface{}, conditions ...interface{}) (int64, error)
	Preload(query string, args ...interface{}) Database
}

type database struct {
	log    Logger
	source *DB
}

var conn = &sync.Mutex{}
var db Database

// Construct for database instance.
func NewDatabase(config *models.Config, log Logger) Database {
	if db == nil {
		conn.Lock()
		defer conn.Unlock()

		if db == nil {
			var source *gorm.DB
			var err error

			// Create new PostgresDB connection
			if source, err = newPostgresDbConnection(config); err != nil {
				log.Fatalf(DB_CONNECTION_ERROR, "error: %v, cause: %v", UNABLE_TO_CONNECT_DB, err)
			}
			db = &database{ // Create singleton instance
				log:    log,
				source: source,
			}
		}
	}

	return db
}

// Implementations for Database interface //

func (db *database) AutoMigrate(models ...interface{}) {
	if err := db.source.AutoMigrate(models...); err != nil {
		db.log.Errorf(DB_MIGRATION_ERROR, "error: failed migrating model %#v, cause: %v", models, err)
	}
}

func (db *database) Create(value interface{}) (int64, error) {
	result := db.source.Create(value)
	return result.RowsAffected, result.Error
}

func (db *database) Find(destination interface{}, conditions ...interface{}) (int64, error) {
	results := db.source.Find(destination, conditions...)
	return results.RowsAffected, results.Error
}

func (db *database) Preload(query string, args ...interface{}) Database {
	return &database{source: db.source.Preload(query, args...)}
}
