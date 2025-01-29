package db

import (
	"github.com/4strodev/4stroblog/site/shared/config"
	"github.com/4strodev/4stroblog/site/shared/db/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var appModels = []any{
	&models.User{},
	&models.Session{},
	&models.Profile{},
}
var dbInstance *gorm.DB

func NewDb(config config.Config) (*gorm.DB, error) {
	if dbInstance != nil {
		return dbInstance, nil
	}

	var err error
	dbInstance, err = gorm.Open(sqlite.Open(config.Db.Sqlite.Path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	for _, model := range appModels {
		err := dbInstance.AutoMigrate(model)
		if err != nil {
			return nil, err
		}
	}

	return dbInstance, nil
}
