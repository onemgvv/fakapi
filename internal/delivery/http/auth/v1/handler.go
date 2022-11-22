package v1

import routing "github.com/qiangxue/fasthttp-routing"

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Init(auth *routing.RouteGroup) {
	v1 := auth.Group("/v1")
	{
		v1.Post("/login", h.Login)
		v1.Post("/register", h.Register)
		v1.Get("/get-code/:id", h.GetCode)
		v1.Post("/restore", h.Restore)
	}
}
