package core

import (
	"sync"

	"github.com/subhankardas/go-microservices/restaurant-service/models"
	"gorm.io/gorm"
)

type DB = gorm.DB

// Wrapper for database functionalities using GORM DB.
type Database interface {
	// Auto-migrate models and creates tables accordingly.
	AutoMigrate(models ...interface{})
	// Inserts new record to DB, return affected rows and error if any.
	Create(value interface{}) (int64, error)
	// Returns all records from DB, matching given conditions, return affected rows and error if any.
	Find(destination interface{}, conditions ...interface{}) (int64, error)
	// Preloads associated tables with given conditions.
	Preload(query string, conditions ...interface{}) Database
	// Update/creates existing or new data in the DB, return affected rows and error if any.
	Save(value interface{}) (int64, error)
	//Finds the first record ordered by primary key, matching given conditions.
	First(value interface{}, conditions ...interface{}) error
	// Deletes value from DB, matching given conditions, return affected rows and error if any.
	Delete(value interface{}, conditions ...interface{}) error
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
			// Create new PostgresDB connection
			source, err := newPostgresDbConnection(config)
			if err != nil {
				log.Fatalf(DB_ERROR, "error: %v, cause: %v", UNABLE_TO_CONNECT_DB, err)
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
		db.log.Errorf(DB_ERROR, "error: failed migrating model %#v, cause: %v", models, err)
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

func (db *database) Preload(query string, conditions ...interface{}) Database {
	return &database{source: db.source.Preload(query, conditions...)}
}

func (db *database) Save(value interface{}) (int64, error) {
	results := db.source.Session(&gorm.Session{FullSaveAssociations: true}).Save(value)
	return results.RowsAffected, results.Error
}

func (db *database) First(value interface{}, conditions ...interface{}) error {
	return db.source.First(value, conditions...).Error
}

func (db *database) Delete(value interface{}, conditions ...interface{}) error {
	return db.source.Delete(value, conditions...).Error
}
