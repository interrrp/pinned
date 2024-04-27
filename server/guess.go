package server

import (
	"log"
	"net/http"
)

func (s *Server) handleGuess(w http.ResponseWriter, r *http.Request) {
	guess, err := readBodyToString(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if s.pinSvc.IsCorrect(guess) {
		newPin := s.pinSvc.Generate()
		log.Printf("PIN has been guessed correctly! New PIN is %s", newPin)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}
