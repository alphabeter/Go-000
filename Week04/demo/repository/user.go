package repository

import (
	"github.com/alphabeter/Go-000/demo/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func ProvideUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

func (u *UserRepository) FindAll() []model.User {
	var users []model.User
	u.DB.Find(&users)
	return users
}

func (u *UserRepository) FindByID(id uint) model.User {
	var user model.User
	u.DB.First(&user, id)
	return user
}