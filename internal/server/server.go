package server

import (
	"net/http"

	"github.com/interrrp/pinned/internal/pin"
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

	s.mux.Handle("/", http.FileServer(http.Dir("./static")))
	s.mux.HandleFunc("POST /", s.handleGuess)

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
