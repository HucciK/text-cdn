package transport

import (
	"fmt"
	"net/http"
	"text-cdn/config"
	"time"
)

type Server struct {
	server *http.Server
}

func NewServer(config config.ServerConfig, h http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:         config.Addr,
			WriteTimeout: time.Duration(config.WriteTimeout) * time.Second,
			ReadTimeout:  time.Duration(config.ReadTimeout) * time.Second,
			Handler:      h,
		},
	}
}

func (s *Server) Start() error {
	if err := s.server.ListenAndServe(); err != nil {
		return fmt.Errorf("error while trying to start server: %v", err)
	}
	return nil
}
