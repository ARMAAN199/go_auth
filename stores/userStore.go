package stores

import (
	"github.com/ARMAAN199/Go_EcomApi/models"
	"gorm.io/gorm"
)

type UserStore interface {
	Login()
	Register(*models.User) (int, error)
	GetUser(string) (*models.User, error)
}

type dbUserStore struct {
	db *gorm.DB
}

func NewDBUserStore(db *gorm.DB) *dbUserStore {
	store := &dbUserStore{
		db: db,
	}
	return store
}

func (store *dbUserStore) Login() {
	var car *models.Cars
	result := store.db.First(&car)

	// Handle the error
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// return nil, result.Error
		} else {
			// return nil, result.Error
		}
	} else {
		// return car, nil
	}
}

func (store *dbUserStore) Register(user *models.User) (int, error) {
	result := store.db.Create(user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

func (store *dbUserStore) GetUser(username string) (*models.User, error) {
	var user *models.User
	result := store.db.First(&user, "username = ?", username)

	// Handle the error
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		} else {
			return nil, result.Error
		}
	} else {
		return user, nil
	}
}
