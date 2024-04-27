package server

import (
	"log/slog"
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
		slog.Info("Regenerated PIN", "newPin", newPin)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}
