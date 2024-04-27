package main

import (
	"net/http"

	"github.com/interrrp/pinned/pin"
)

type Server struct {
	pinSvc pin.PINService
	mux    *http.ServeMux
}

func NewServer(pinSvc pin.PINService) *Server {
	s := &Server{
		pinSvc: pinSvc,
		mux:    http.NewServeMux(),
	}

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
