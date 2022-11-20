package users

import (
	apiDto "github.com/onemgvv/fakapi/internal/delivery/http/api/v1/dto"
	authDto "github.com/onemgvv/fakapi/internal/delivery/http/auth/v1/dto"
	"github.com/onemgvv/fakapi/internal/domain"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) Create(dto authDto.RegisterUserDto) *User {
	return NewUser(dto)
}

func (us *UserService) All() map[int]User {
	return domain.Users
}

func (us *UserService) GetById(ID int) User {
	return domain.Users[ID]
}

func (us *UserService) Update(ID int, dto apiDto.UpdateUserDto) {
	user := domain.Users[ID]

	user.Age = dto.Age
	user.GivenName = dto.GivenName
	user.FamilyName = dto.FamilyName
	user.Nickname = dto.Nickname
	user.Phone = dto.Phone

	domain.Users[ID] = user
}

func (us *UserService) ChangePassword(ID int, OldPassword, NewPassword string) {
	user := domain.Users[ID]
	if !user.ComparePassword(OldPassword) {
		return
	}

	user.NewPassword(NewPassword)
}

func (us *UserService) Delete(ID int) {
	delete(domain.Users, ID)
}
