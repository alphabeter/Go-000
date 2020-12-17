package api

import (
	"net/http"

	"github.com/alphabeter/Go-000/demo/mapper"
	"github.com/alphabeter/Go-000/demo/service"
	"github.com/gin-gonic/gin"
)

// UserAPI struct.
type UserAPI struct {
	UserService service.UserService
}

// ProvideUserAPI func.
func ProvideUserAPI(u service.UserService) UserAPI {
	return UserAPI{UserService: u}
}

// FindAll method for UserAPI.
func (u *UserAPI) FindAll(c *gin.Context) {
	users := u.UserService.FindAll()
	c.JSON(http.StatusOK, gin.H{"users": mapper.ToUserDTOs(users)})
}
