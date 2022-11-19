package users

import (
	"github.com/onemgvv/fakapi/internal/domain"
	"strings"
	"time"
)

type User struct {
	Id         int    `json:"id"`
	FamilyName string `json:"familyName"`
	GivenName  string `json:"givenName"`
	Nickname   string `json:"nickname"`
	Age        int    `json:"age"`
	Phone      string `json:"phone"`
	password   string
	RegDate    time.Time `json:"regDate"`
}

func (u *User) ComparePassword(password string) bool {
	return strings.EqualFold(u.password, password)
}

func NewUser(fName, gName, nick, password, phone string, age int) *User {
	user := User{
		Id:         len(domain.Users) + 1,
		FamilyName: fName,
		GivenName:  gName,
		Nickname:   nick,
		password:   password,
		Phone:      phone,
		Age:        age,
		RegDate:    time.Now(),
	}
	
	domain.Users = append(domain.Users, user)

	return &user
}
