package service

import (
	"fmt"

	"playground.io/another-pet-store/dto"
)

type LoginService interface {
	LoginUser(user *dto.User) bool
}

type localUser dto.User

func DbLoginService() LoginService {
	return &localUser{}
}

func (out *localUser) LoginUser(user *dto.User) bool {
	var dbUser = FindUserByUsername(user.Username)
	var isIt = CheckPasswordHash(user.Password, dbUser.Hash)
	fmt.Print("Login and pass are ", isIt)
	return isIt
}
