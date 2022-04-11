package service

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
	"playground.io/another-pet-store/db"
	"playground.io/another-pet-store/dto"
	"playground.io/another-pet-store/model"
)

type UserService interface {
	RegisterUser(user *dto.User)
	FindUserByUsername(username string) *model.User
}

type SimpleUserService struct {
}

func NewUserService() UserService {
	return &SimpleUserService{}
}

func (s *SimpleUserService) FindUserByUsername(username string) *model.User {
	return db.FindUser(username)
}

func (s *SimpleUserService) RegisterUser(userDto *dto.User) {
	fmt.Println(userDto.Password)
	hash, err := hashPassword(userDto.Password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to generate hash: %v\n", err)
		os.Exit(1)
	}

	var user = model.User{Username: userDto.Username, Hash: hash, Email: userDto.Email}
	db.AddUser(&user)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
