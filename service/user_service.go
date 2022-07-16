package service

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"playground.io/another-pet-store/db"
	"playground.io/another-pet-store/dto"
	"playground.io/another-pet-store/model"
)

type UserService interface {
	RegisterUser(user *dto.User) error
	FindUserByUsername(username string) (*model.User, error)
}

type SimpleUserService struct {
	userRepository db.UserRepository
}

func NewUserService(repository db.UserRepository) UserService {
	return &SimpleUserService{userRepository: repository}
}

func (s *SimpleUserService) FindUserByUsername(username string) (*model.User, error) {
	user, err := s.userRepository.FindUser(username)
	if err != nil {
		return nil, fmt.Errorf("user '%s' not found: %w", username, err)
	}
	return user, nil
}

func (s *SimpleUserService) RegisterUser(userDto *dto.User) error {
	hash, err := hashPassword(userDto.Password)
	if err != nil {
		return fmt.Errorf("failed to generate hash: %v", err)
	}

	var user = model.User{Username: userDto.Username, Hash: hash, Email: userDto.Email}
	s.userRepository.AddUser(&user)
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
