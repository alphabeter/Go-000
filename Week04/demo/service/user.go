package service

import (
	"github.com/alphabeter/Go-000/demo/model"
	"github.com/alphabeter/Go-000/demo/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func ProvideUserService(u repository.UserRepository) UserService {
	return UserService{UserRepository: u}
}

func (u *UserService) FindAll() []model.User {
	return u.UserRepository.FindAll()
}

func (u *UserService) FindByID(id uint) model.User {
	return u.UserRepository.FindByID(id)
}