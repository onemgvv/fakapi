package server

import (
	"github.com/valyala/fasthttp"
)

type Server struct {
	httpServer *fasthttp.Server
}

func NewServer(handler fasthttp.RequestHandler) *Server {
	return &Server{
		httpServer: &fasthttp.Server{
			Handler: handler,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe(":7888")
}

func (s *Server) Stop() error {
	return s.httpServer.Shutdown()
}
