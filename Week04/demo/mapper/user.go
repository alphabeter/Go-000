package mapper

import (
	"github.com/alphabeter/Go-000/demo/dto"
	"github.com/alphabeter/Go-000/demo/model"
)

// ToUser func.
func ToUser(dto dto.UserDTO) model.User {
	return model.User{Name: dto.Name, Age: dto.Age}
}

// ToUserDTO func.
func ToUserDTO(u model.User) dto.UserDTO {
	return dto.UserDTO{ID: u.ID, Name: u.Name, Age: u.Age}
}

// ToUserDTOs func.
func ToUserDTOs(u []model.User) []dto.UserDTO {
	dtos := make([]dto.UserDTO, len(u))
	for i, item := range u {
		dtos[i] = ToUserDTO(item)
	}
	return dtos
}
