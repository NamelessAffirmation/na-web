package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	accessTokenExpirationInMinutes         = 60
	verificationTokenExpirationInHours     = 3
	loginWithEmailTokenExpirationInMinutes = 10
	refreshTokenExpirationInHours          = 160
)

type TokenService struct {
	JWTSecretKey string
}

func NewTokenService(jwtSecretKey string) *TokenService {
	return &TokenService{
		JWTSecretKey: jwtSecretKey,
	}
}

func (s *TokenService) GenerateAccessToken(id int) (*string, error) {

	expirationTime := time.Now().Add(accessTokenExpirationInMinutes * time.Minute).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512,
		jwt.MapClaims{
			"iss":  "https://myrankedlist.com",
			"sub":  "myrankedlistapi_access_token",
			"exp":  expirationTime,
			"user": id,
		})
	signedToken, err := token.SignedString([]byte(s.JWTSecretKey))
	if err != nil {
		return nil, err
	}
	return &signedToken, nil
}

func (s *TokenService) GenerateLoginWithEmailToken(id int) (*string, error) {

	expirationTime := time.Now().Add(loginWithEmailTokenExpirationInMinutes * time.Minute).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512,
		jwt.MapClaims{
			"iss":  "https://myrankedlist.com",
			"sub":  "myrankedlistapi_loginwithemail_token",
			"exp":  expirationTime,
			"user": id,
		})
	signedToken, err := token.SignedString([]byte(s.JWTSecretKey))
	if err != nil {
		return nil, err
	}
	return &signedToken, nil
}

func (s *TokenService) GenerateVerificationToken(id int) (*string, error) {

	expirationTime := time.Now().Add(verificationTokenExpirationInHours * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512,
		jwt.MapClaims{
			"iss":  "https://myrankedlist.com",
			"sub":  "myrankedlistapi_verification_token",
			"exp":  expirationTime,
			"user": id,
		})
	signedToken, err := token.SignedString([]byte(s.JWTSecretKey))
	if err != nil {
		return nil, err
	}
	return &signedToken, nil
}

func (s *TokenService) GenerateRefreshToken(id int) (*string, error) {

	expirationTime := time.Now().Add(refreshTokenExpirationInHours * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512,
		jwt.MapClaims{
			"iss":  "https://myrankedlist.com",
			"sub":  "myrankedlistapi_refresh_token",
			"exp":  expirationTime,
			"user": id,
		})
	signedToken, err := token.SignedString([]byte(s.JWTSecretKey))
	if err != nil {
		return nil, err
	}
	return &signedToken, nil
}

func (s *TokenService) ValidateToken(tokenToVerify *string) (*int, error) {

	parsedToken, err := jwt.Parse(*tokenToVerify, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(s.JWTSecretKey), nil
	})
	if err != nil {
		return nil, errors.New("JWT could not be validated")
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
		userId := int(claims["user"].(float64))
		return &userId, nil
	} else {
		return nil, errors.New("JWT claims could not be validated")
	}
}
