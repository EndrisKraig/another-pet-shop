package middleware

import (
	"fmt"
	"net/http"

	"playground.io/another-pet-store/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeader     = "Authorization"
	AuthorizationPayloadKey = "Authorization_payload"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		token, err := service.NewJWTService().ValidateToken(authHeader)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			c.Set(AuthorizationPayloadKey, claims)
			return
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

func GetClaims(c *gin.Context) (jwt.MapClaims, error) {
	payload, ok := c.Get(AuthorizationPayloadKey)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return jwt.MapClaims{}, fmt.Errorf("no auth payload")
	}

	return payload.(jwt.MapClaims), nil
}
