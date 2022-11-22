package domain

import (
	apiDto "github.com/onemgvv/fakapi/internal/delivery/http/api/v1/dto"
	authDto "github.com/onemgvv/fakapi/internal/delivery/http/auth/v1/dto"
	"github.com/onemgvv/fakapi/internal/domain/auth"
	"github.com/onemgvv/fakapi/internal/domain/users"
)

var (
	Users = make(map[int]users.User, 20)
)

type Authorization interface {
	Login()
	Register()
	GetCode()
	Restore()
}

type UserService interface {
	Create(dto authDto.RegisterUserDto) *users.User
	All() map[int]users.User
	GetById(ID int) users.User
	Update(ID int, dto apiDto.UpdateUserDto)
	ChangePassword(ID int, OldPassword, NewPassword string)
	Delete(ID int)
}

type Service struct {
	Authorization
	UserService
}

func NewService() *Service {
	return &Service{
		Authorization: auth.NewAuthorization(),
		UserService:   users.NewUserService(),
	}
}
