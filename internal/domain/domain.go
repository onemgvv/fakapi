package domain

import (
	apiDto "github.com/onemgvv/fakapi/internal/delivery/http/api/v1/dto"
	authDto "github.com/onemgvv/fakapi/internal/delivery/http/auth/v1/dto"
	"github.com/onemgvv/fakapi/internal/domain/users"
)

var (
	Users map[int]users.User = make(map[int]users.User, 20)
)

type UserService interface {
	Create(dto authDto.RegisterUserDto) *users.User
	All() map[int]users.User
	GetById(ID int) users.User
	Update(ID int, dto apiDto.UpdateUserDto)
	ChangePassword(ID int, OldPassword, NewPassword string)
	Delete(ID int)
}

type Service struct {
	UserService
}

func NewService() *Service {
	return &Service{
		UserService: users.NewUserService(),
	}
}
