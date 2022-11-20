package users

import (
	"github.com/onemgvv/fakapi/internal/delivery/http/auth/v1/dto"
	"github.com/onemgvv/fakapi/internal/domain"
	"strings"
	"time"
)

type User struct {
	Id         int    `json:"id"`
	FamilyName string `json:"familyName"`
	GivenName  string `json:"givenName"`
	Nickname   string `json:"nickname"`
	Age        uint   `json:"age"`
	Phone      string `json:"phone"`
	password   string
	RegDate    time.Time `json:"regDate"`
}

func (u *User) ComparePassword(password string) bool {
	return strings.EqualFold(u.password, password)
}

func (u *User) NewPassword(password string) {
	u.password = password
}

func NewUser(dto dto.RegisterUserDto) *User {
	user := User{
		Id:         len(domain.Users) + 1,
		FamilyName: dto.FamilyName,
		GivenName:  dto.GivenName,
		Nickname:   dto.Nickname,
		password:   dto.Password,
		Phone:      dto.Phone,
		Age:        dto.Age,
		RegDate:    time.Now(),
	}

	domain.Users[user.Id] = user

	return &user
}
