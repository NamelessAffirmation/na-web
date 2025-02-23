package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/snowlynxsoftware/na-api/server/middleware"
	"github.com/snowlynxsoftware/na-api/server/models"
	"github.com/snowlynxsoftware/na-api/server/services"
)

type AuthController struct {
	AuthMiddleware *middleware.AuthMiddleware
	UserService    *services.UserService
	AuthService    *services.AuthService
}

func NewAuthController(authMiddleware *middleware.AuthMiddleware, userService *services.UserService, authService *services.AuthService) *AuthController {
	return &AuthController{
		AuthMiddleware: authMiddleware,
		UserService:    userService,
		AuthService:    authService,
	}
}

func (c *AuthController) MapController() *chi.Mux {
	router := chi.NewRouter()
	// Public Routes
	router.Post("/login", c.login)
	router.Post("/register", c.register)
	router.Get("/verify", c.verify)
	router.Post("/send-login-email", c.sendLoginEmail)
	router.Get("/login-with-email", c.loginWithEmail)

	// Protected Routes
	router.Get("/token", c.tokenInfo)
	return router
}

func (c *AuthController) login(w http.ResponseWriter, r *http.Request) {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header is required", http.StatusUnauthorized)
		return
	}

	response, err := c.AuthService.Login(&authHeader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	returnStr, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "failed to create response", http.StatusInternalServerError)
		return
	}

	fmt.Println(response.AccessToken)

	http.SetCookie(w, &http.Cookie{
		Name:     "mrl_access_token",
		Value:    response.AccessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // Set to true if using HTTPS
		SameSite: http.SameSiteLaxMode,
		MaxAge:   59 * 60,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(returnStr)
}

func (c *AuthController) sendLoginEmail(w http.ResponseWriter, r *http.Request) {
	var userCreateDTO models.UserCreateDTO

	err := json.NewDecoder(r.Body).Decode(&userCreateDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if userCreateDTO.Email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	userEntity, err := c.AuthService.SendLoginEmail(strings.ToLower(userCreateDTO.Email))
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "an error occurred when attempting to send the login email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("user (%v) email was sent.", userEntity.Email)))
}

func (c *AuthController) register(w http.ResponseWriter, r *http.Request) {
	var userCreateDTO models.UserCreateDTO

	err := json.NewDecoder(r.Body).Decode(&userCreateDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if userCreateDTO.Username == "" || userCreateDTO.Password == "" || userCreateDTO.Email == "" {
		http.Error(w, "Email, Username, and Password are required", http.StatusBadRequest)
		return
	}

	userEntity, err := c.AuthService.RegisterNewUser(&userCreateDTO)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "an error occurred when attempting to register your user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("user (%v) was created. check your email for verification.", userEntity.Email)))
}

func (c *AuthController) loginWithEmail(w http.ResponseWriter, r *http.Request) {

	verificationToken := r.URL.Query().Get("token")

	userId, err := c.AuthService.VerifyNewUser(&verificationToken)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response, err := c.AuthService.LoginWithEmailLink(userId)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "mrl_access_token",
		Value:    response.AccessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // Set to true if using HTTPS
		SameSite: http.SameSiteLaxMode,
		MaxAge:   59 * 60,
	})

	w.Header().Set("Content-Type", "application/json")
	http.Redirect(w, r, "/lists", http.StatusSeeOther)

}

func (c *AuthController) verify(w http.ResponseWriter, r *http.Request) {

	verificationToken := r.URL.Query().Get("token")

	_, err := c.AuthService.VerifyNewUser(&verificationToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user was verified successfully. you can now login"))

}

func (c *AuthController) tokenInfo(w http.ResponseWriter, r *http.Request) {

	userContext, err := c.AuthMiddleware.Authorize(r, nil)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "an error occurred when attempting to get token info", http.StatusUnauthorized)
		return
	}

	returnStr, err := json.Marshal(userContext)
	if err != nil {
		http.Error(w, "failed to create response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(returnStr)

}
