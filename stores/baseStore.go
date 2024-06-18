package stores

import (
	"github.com/ARMAAN199/Go_EcomApi/models"
	"gorm.io/gorm"
)

type BaseStore interface {
	GetCar() (*models.Cars, error)
}

type dbBaseStore struct {
	db *gorm.DB
}

func NewDBBaseStore(db *gorm.DB) *dbBaseStore {
	store := &dbBaseStore{
		db: db,
	}
	return store
}

func (store *dbBaseStore) GetCar() (*models.Cars, error) {
	var car *models.Cars
	result := store.db.First(&car)

	// Handle the error
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		} else {
			return nil, result.Error
		}
	} else {
		return car, nil
	}
}
