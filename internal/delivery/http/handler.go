package http

import (
	"github.com/gin-gonic/gin"
	"github.com/onemgvv/fakapi/internal/delivery/http/middleware"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	gin.SetMode(gin.DebugMode) // Development Mode
	//gin.SetMode(gin.ReleaseMode) // Production Mode

	router.Use(
		gin.Logger(),
		gin.Recovery(),
		middleware.Cors,
	)

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "%s\n", "pong")
	})

	auth := router.Group("/auth")
	{
		auth.POST("/login")
		auth.POST("/register")
		auth.GET("/get-code/:id")
		auth.POST("/restore")
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

	return router
}
