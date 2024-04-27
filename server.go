package main

import (
	"io"
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

	s.mux.HandleFunc("POST /", s.handleGuess)

	return s
}

func (s *Server) handleGuess(w http.ResponseWriter, r *http.Request) {
	guess, err := readBodyToString(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if s.pinSvc.IsCorrect(guess) {
		w.WriteHeader(http.StatusOK)
		s.pinSvc.Generate()
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func readBodyToString(r *http.Request) (string, error) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
