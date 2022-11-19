package http

import (
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

	auth := router.Group("/auth")
	{
		auth.Post("/login")
		auth.Post("/register")
		auth.Get("/get-code/<id>")
		auth.Post("/restore")
	}

	api := router.Group("/api")
	{
		api.Group("/users")
		api.Group("/countries")
		api.Group("/cities")
		api.Group("/addresses")
		api.Group("/posts")
		api.Group("/products")
		api.Group("/weather")
		api.Group("/music")
		api.Group("/books")
		api.Group("/services")
		api.Group("/images")
	}

	return router.HandleRequest
}
