package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type SettingsController struct {
}

func NewSettingsController() *SettingsController {
	return &SettingsController{}
}

func (s *SettingsController) MapController() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/client-data", func(w http.ResponseWriter, r *http.Request) {
		returnStr, err := json.Marshal([]byte("{\"test\": \"test\"}"))
		if err != nil {
			http.Error(w, "failed to create response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(returnStr)
	})
	return r
}
