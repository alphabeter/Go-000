package repository

import (
	"github.com/alphabeter/Go-000/demo/model"
	"gorm.io/gorm"
)

// UserRepository struct.
type UserRepository struct {
	DB *gorm.DB
}

// ProvideUserRepository returns a UserRepository.
func ProvideUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

// FindAll method for UserRepository.
func (u *UserRepository) FindAll() []model.User {
	var users []model.User
	u.DB.Find(&users)
	return users
}

// FindByID method for UserRepository.
func (u *UserRepository) FindByID(id uint) model.User {
	var user model.User
	u.DB.First(&user, id)
	return user
}
