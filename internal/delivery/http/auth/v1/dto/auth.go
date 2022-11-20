package dto

type RegisterUserDto struct {
	FamilyName string `json:"familyName"`
	GivenName  string `json:"givenName"`
	Nickname   string `json:"nickname"`
	Phone      string `json:"phone"`
	Age        uint   `json:"age"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type LoginUserDto struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
