package api

import (
	"github.com/alphabeter/Go-000/demo/mapper"
	"github.com/alphabeter/Go-000/demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserAPI struct {
	UserService service.UserService
}

func ProvideUserAPI(u service.UserService) UserAPI {
	return UserAPI{UserService: u}
}

func (u *UserAPI) FindAll(c *gin.Context) {
	users := u.UserService.FindAll()
	c.JSON(http.StatusOK, gin.H{"users": mapper.ToUserDTOs(users)})
}