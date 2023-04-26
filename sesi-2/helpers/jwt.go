package helpers

import "github.com/go-kit/kit/auth/jwt"

var secretKey = "ThisistheK3y"

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}
}
