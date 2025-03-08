package service

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct{}

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

func (ps *PasswordService) HashPassword(password []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	log.Println(hash)
	return hash, err
}

func (ps *PasswordService) CheckPasswordHash(password, hash []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}
