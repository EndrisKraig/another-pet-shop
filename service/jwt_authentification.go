package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(email string, isUser bool, userId, profileId int) string
	ValidateToken(token string) (*jwt.Token, error)
}

type AuthCustomClaims struct {
	Name      string `json:"name"`
	User      bool   `json:"user"`
	UserId    int    `json:"userId"`
	ProfileId int    `json:"profileId"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issuer:    "Admin",
	}
}

func (service *jwtServices) GenerateToken(email string, isUser bool, userId, profileId int) string {
	claims := &AuthCustomClaims{
		email,
		isUser,
		userId,
		profileId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token %s", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})

}
