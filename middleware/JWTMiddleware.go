package middleware

import (
	"fmt"
	"net/http"

	"playground.io/another-pet-store/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		Me(c)
	}
}

func Me(c *gin.Context) (jwt.MapClaims, error) {
	authHeader := c.GetHeader("Authorization")
	fmt.Println(authHeader)
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return nil, fmt.Errorf("No auth header!")
	}
	token, err := service.JWTAuthService().ValidateToken(authHeader)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return nil, fmt.Errorf("No auth header!")
	}
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		fmt.Println(claims)
		return claims, nil
	} else {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return nil, err
	}
}
