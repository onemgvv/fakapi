package http

import (
	apiV1 "github.com/onemgvv/fakapi/internal/delivery/http/api/v1"
	authV1 "github.com/onemgvv/fakapi/internal/delivery/http/auth/v1"
	"github.com/onemgvv/fakapi/internal/delivery/http/middleware"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() fasthttp.RequestHandler {
	router := routing.New()

	router.Use(
		middleware.Cors,
	)

	router.Get("/ping", func(c *routing.Context) error {
		c.SuccessString("application/json", "pong")
		return nil
	})

	h.InitAuth(router)
	h.InitAPI(router)

	return router.HandleRequest
}

func (h *Handler) InitAuth(router *routing.Router) {
	handlerV1 := authV1.NewAuthHandler()
	auth := router.Group("/auth")
	{
		handlerV1.Init(auth)
	}
}

func (h *Handler) InitAPI(router *routing.Router) {
	handlerV1 := apiV1.NewAPIHandler()
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
