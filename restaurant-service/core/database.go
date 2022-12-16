package core

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/postgres"
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

func NewDatabase(log Logger) Database {
	if db == nil {
		conn.Lock()
		defer conn.Unlock()

		if db == nil {
			conn, err := newDBConnection()
			if err != nil {
				log.Fatalf(DB_CONNECTION_ERROR, "Could not connect to DB, error: %v", err)
			}
			db = &database{
				log:    log,
				source: conn,
			}
		}
	}

	return db
}

func newDBConnection() (*DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, username, password, databaseName, port)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// Implementations for Database interface //

func (db *database) AutoMigrate(models ...interface{}) {
	if err := db.source.AutoMigrate(models...); err != nil {
		db.log.Errorf(DB_MIGRATION_ERROR, "Error migrating model %#v, error: %v", models, err)
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
