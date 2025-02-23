package services

import (
	"crypto/sha512"

	"golang.org/x/crypto/bcrypt"
)

var BCRYPT_COST = 13

type CryptoService struct {
	Pepper string
}

func NewCryptoService(pepper string) *CryptoService {
	return &CryptoService{
		Pepper: pepper,
	}
}

func (s *CryptoService) HashPassword(password string) (*string, error) {

	var bytesToHash = pepperAndHashPassword(password, s.Pepper)

	var hashBytes, err = bcrypt.GenerateFromPassword(bytesToHash, BCRYPT_COST)
	if err != nil {
		return nil, err
	}
	var hashedString = string(hashBytes)

	return &hashedString, nil
}

func (s *CryptoService) ValidatePassword(password string, hash string) (bool, error) {

	var bytesToHash = pepperAndHashPassword(password, s.Pepper)

	var err = bcrypt.CompareHashAndPassword([]byte(hash), bytesToHash)
	if err != nil {
		return false, err
	}
	return true, nil
}

func pepperAndHashPassword(password string, pepper string) []byte {
	var shaSumBytes = sha512.Sum512([]byte(password + pepper))
	return shaSumBytes[:]
}
