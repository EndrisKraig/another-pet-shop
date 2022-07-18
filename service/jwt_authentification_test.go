package service_test

import (
	"testing"

	"playground.io/another-pet-store/service"
)

func TestGenerateToken(t *testing.T) {
	jwtService := service.NewJWTService()

	token := jwtService.GenerateToken("sparrow", true, 1, 1)
	tokenWrapper, err := jwtService.ValidateToken(token)
	assertNoError(err, t)
	if tokenWrapper == nil {
		t.Errorf("Expected token, got nil")
	}
}
