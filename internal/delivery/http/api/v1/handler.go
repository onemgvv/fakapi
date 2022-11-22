package v1

import routing "github.com/qiangxue/fasthttp-routing"

type APIHandler struct{}

func NewAPIHandler() *APIHandler {
	return &APIHandler{}
}

func (h *APIHandler) Init(api *routing.RouteGroup) {
	v1 := api.Group("/v1")
	users := v1.Group("/users")
	{
		users.Patch("/profile/:id", h.UpdateProfile)
		users.Patch("/access/:id", h.UpdatePassword)
		users.Delete("/:id", h.Delete)
		users.Get("/", h.All)
		users.Get("/:id", h.GetOne)
	}
	v1.Group("/countries")
	v1.Group("/cities")
	v1.Group("/addresses")
	v1.Group("/posts")
	v1.Group("/products")
	v1.Group("/weather")
	v1.Group("/music")
	v1.Group("/books")
	v1.Group("/services")
	v1.Group("/images")
}
