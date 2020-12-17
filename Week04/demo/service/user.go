package service

import (
	"github.com/alphabeter/Go-000/demo/model"
	"github.com/alphabeter/Go-000/demo/repository"
)

// UserService struct.
type UserService struct {
	UserRepository repository.UserRepository
}

// ProvideUserService returns a UserService.
func ProvideUserService(u repository.UserRepository) UserService {
	return UserService{UserRepository: u}
}

// FindAll method for UserService.
func (u *UserService) FindAll() []model.User {
	return u.UserRepository.FindAll()
}

// FindByID method for UserService.
func (u *UserService) FindByID(id uint) model.User {
	return u.UserRepository.FindByID(id)
}
