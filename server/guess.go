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

	correct, err := s.pinSvc.IsCorrect(guess)
	if err != nil {
		slog.Error("Failed to judge guess", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if correct {
		newPin, err := s.pinSvc.Generate()
		if err != nil {
			slog.Error("Failed to regenerate new PIN", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		slog.Info("Regenerated PIN", "newPin", newPin)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}
