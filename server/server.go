package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Server struct {
	srv http.Server
}

func New(r *mux.Router) Server {
	return Server{
		srv: http.Server{
			Handler: r,
		},
	}
}

func (s *Server) Setup() {
	s.srv.Addr = "127.0.0.1:8080"
	s.srv.WriteTimeout = 15 * time.Second
	s.srv.ReadTimeout = 15 * time.Second
}

func (s *Server) Run() {
	s.srv.ListenAndServe()
}
