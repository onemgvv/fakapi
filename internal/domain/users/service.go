package users

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) Create() *User {
	return nil
}
