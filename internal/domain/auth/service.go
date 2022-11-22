package auth

type Authorization struct{}

func NewAuthorization() *Authorization {
	return &Authorization{}
}

func (as *Authorization) Login()    {}
func (as *Authorization) Register() {}
func (as *Authorization) GetCode()  {}
func (as *Authorization) Restore()  {}
