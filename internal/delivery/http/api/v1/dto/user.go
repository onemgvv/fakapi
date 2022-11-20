package dto

type UpdateUserDto struct {
	FamilyName string `json:"familyName"`
	GivenName  string `json:"givenName"`
	Phone      string `json:"phone"`
	Age        uint   `json:"age"`
	Nickname   string `json:"nickname"`
}
