package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/snowlynxsoftware/na-api/server/services"
)

type AuthMiddleware struct {
	UserService *services.UserService
}

func NewAuthMiddleware(userService *services.UserService) *AuthMiddleware {
	return &AuthMiddleware{
		UserService: userService,
	}
}

// If a request is authorized, it ill return this context to the controller
// so that information from the user can be used as an immutable object.
type AuthorizedUserContext struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (m *AuthMiddleware) Authorize(r *http.Request, requiredUserTypeKeys []string) (*AuthorizedUserContext, error) {

	cookie, err := r.Cookie("mrl_access_token")
	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("access token not found in request")
	}

	userId, err := m.UserService.TokenService.ValidateToken(&cookie.Value)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	userEntity, err := m.UserService.GetUserById(*userId)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	if userEntity.IsArchived {
		return nil, errors.New("user is archived")
	}

	if !userEntity.IsVerified {
		return nil, errors.New("user is not verified")
	}

	if userEntity.IsBanned {
		return nil, errors.New("user is banned. Reason - " + *userEntity.BanReason)
	}

	// TODO: Check User Type here.

	return &AuthorizedUserContext{
		Id:       int(userEntity.ID),
		Email:    userEntity.Email,
		Username: userEntity.Username,
	}, nil

}
