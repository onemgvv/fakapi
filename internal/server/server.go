package server

import (
	"github.com/onemgvv/fakapi/internal/config"
	"github.com/valyala/fasthttp"
)

type Server struct {
	httpServer *fasthttp.Server
}

func NewServer(cfg *config.Config, handler fasthttp.RequestHandler) *Server {
	return &Server{
		httpServer: &fasthttp.Server{
			Handler:      handler,
			ReadTimeout:  cfg.HTTP.Timeout.Read,
			WriteTimeout: cfg.HTTP.Timeout.Write,
		},
	}
}

func (s *Server) Run(port string) error {
	return s.httpServer.ListenAndServe(port)
}

func (s *Server) Stop() error {
	return s.httpServer.Shutdown()
}
