package service_test

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"playground.io/another-pet-store/dto"
	"playground.io/another-pet-store/model"
	"playground.io/another-pet-store/service"
	"testing"
)

type StubUserService struct {
	users map[string]dto.User
}

func (s *StubUserService) RegisterUser(user *dto.User) error {
	s.users[user.Username] = *user
	return nil
}

func (s *StubUserService) FindUserByUsername(username string) (*model.User, error) {
	user, ok := s.users[username]
	if !ok {
		return nil, fmt.Errorf("no user with username %s", username)
	}
	return &model.User{ID: user.ID, Username: user.Username}, nil
}

type StubJWTService struct {
}

func (s StubJWTService) GenerateToken(email string, isUser bool, userId, profileId int) string {
	return "1"
}

func (s StubJWTService) ValidateToken(token string) (*jwt.Token, error) {
	return &jwt.Token{}, nil
}

func TestLoginService(t *testing.T) {
	loginService := service.NewLoginService(&StubUserService{users: map[string]dto.User{}}, &StubProfileService{}, StubJWTService{})

	loginService.NewUser(&dto.User{Username: "Algos"})

	token, err := loginService.LoginUser(&dto.User{Username: "Algos"})
	assertNoError(err, t)
	if token == "" {
		t.Fatalf("Expected real token, not an empty string")
	}

	_, err = loginService.LoginUser(&dto.User{})
	assertError(err, t)
}
