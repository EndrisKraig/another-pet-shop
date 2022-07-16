package service_test

import (
	"fmt"
	"testing"

	"playground.io/another-pet-store/model"
	"playground.io/another-pet-store/service"
)

type SpyUserRepository struct {
	users map[string]model.User
}

func (r *SpyUserRepository) AddUser(user *model.User) error {
	return nil
}

func (r *SpyUserRepository) FindUser(name string) (*model.User, error) {
	user, ok := r.users[name]
	if !ok {
		return nil, fmt.Errorf("No user")
	}
	return &user, nil
}

func TestFindUser(t *testing.T) {
	users := map[string]model.User{
		"Avarosa":   model.User{},
		"Barbadosa": model.User{},
		"Caren":     model.User{},
	}
	userService := service.NewUserService(&SpyUserRepository{users: users})

	t.Run("get existing user", func(t *testing.T) {
		user, err := userService.FindUserByUsername("Avarosa")
		assertNoError(err, t)
		if user == nil {
			t.Errorf("wanted user but didn't get one")
		}
	})

	t.Run("no such user", func(t *testing.T) {
		_, err := userService.FindUserByUsername("Garandel")
		assertError(err, t)
	})

}

func assertError(err error, t *testing.T) {
	t.Helper()
	if err == nil {
		t.Errorf("Wanted a error but didn't get any")
	}
}

func assertNoError(err error, t *testing.T) {
	t.Helper()
	if err != nil {
		t.Errorf("Got an error but didn't want any")
	}
}
