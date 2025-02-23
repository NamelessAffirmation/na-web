package server

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	mid "github.com/go-chi/chi/v5/middleware"
	"github.com/snowlynxsoftware/na-api/config"
	"github.com/snowlynxsoftware/na-api/server/controllers"
	"github.com/snowlynxsoftware/na-api/server/database"
	"github.com/snowlynxsoftware/na-api/server/database/repositories"
	"github.com/snowlynxsoftware/na-api/server/middleware"
	"github.com/snowlynxsoftware/na-api/server/services"
)

//go:embed templates/*
var templatesFS embed.FS

//go:embed public/*
var publicFS embed.FS

type AppServer struct {
	AppConfig *config.AppConfig
	Router    *chi.Mux
	DB        *database.AppDataSource
}

func NewAppServer() *AppServer {

	r := chi.NewRouter()
	r.Use(mid.Logger)

	return &AppServer{
		AppConfig: config.NewAppConfig(),
		Router:    r,
	}
}

func (s *AppServer) Start() {

	// Connect to DB
	s.DB = database.NewAppDataSource()
	s.DB.Connect(s.AppConfig.DBConnectionString)

	// Configure Repositories
	userRepository := repositories.NewUserRepository(s.DB)

	// Configure Services
	emailService := services.NewEmailService(s.AppConfig.SendgridAPIKey)
	cryptoService := services.NewCryptoService(s.AppConfig.HashPepper)
	tokenService := services.NewTokenService(s.AppConfig.JWTSecret)
	authService := services.NewAuthService(userRepository, tokenService, cryptoService, emailService)
	userService := services.NewUserService(userRepository, tokenService)

	// Configure Middleware
	authMiddleware := middleware.NewAuthMiddleware(userService)

	// Configure Controllers
	s.Router.Mount("/health", controllers.NewHealthController().MapController())
	s.Router.Mount("/api/settings", controllers.NewSettingsController().MapController())
	s.Router.Mount("/api/users", controllers.NewUserController(userService).MapController())
	s.Router.Mount("/api/auth", controllers.NewAuthController(authMiddleware, userService, authService).MapController())
	s.Router.Mount("/", controllers.NewUIController(&templatesFS, authMiddleware).MapController())

	// Configure File Server
	s.Router.Handle("/public/*", http.FileServer(http.FS(publicFS)))

	fmt.Printf("Server starting on port %s\n", s.AppConfig.AppPort)
	log.Fatal(http.ListenAndServe(":"+s.AppConfig.AppPort, s.Router))
}
