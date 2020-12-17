//+build wireinject

package main

import (
	"github.com/alphabeter/Go-000/demo/api"
	"github.com/alphabeter/Go-000/demo/repository"
	"github.com/alphabeter/Go-000/demo/service"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// InitUserAPI returns a UserAPI.
func InitUserAPI(db *gorm.DB) api.UserAPI {
	wire.Build(repository.ProvideUserRepository, service.ProvideUserService, api.ProvideUserAPI)
	return api.UserAPI{}
}
