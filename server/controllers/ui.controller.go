package controllers

import (
	"embed"
	"fmt"
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/snowlynxsoftware/na-api/server/middleware"
)

type UIController struct {
	templatesFS    *embed.FS
	AuthMiddleware *middleware.AuthMiddleware
}

func NewUIController(templates *embed.FS, authMiddleware *middleware.AuthMiddleware) *UIController {
	return &UIController{
		templatesFS:    templates,
		AuthMiddleware: authMiddleware,
	}
}

func (s *UIController) MapController() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", s.loginPage)
	r.Get("/login", s.loginPage)
	r.Get("/register", s.registerPage)
	return r
}

func (s *UIController) registerPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFS(s.templatesFS, "templates/components/navbar.html", "templates/components/follow_us.html", "templates/register.html", "templates/base_layout.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing template: %v", err), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "base", nil)
}

func (s *UIController) loginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").ParseFS(s.templatesFS, "templates/components/navbar.html", "templates/components/follow_us.html", "templates/login.html", "templates/base_layout.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing template: %v", err), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "base", nil)
}
