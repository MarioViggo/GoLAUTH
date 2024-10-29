package dados

import (
	"log"

	"github.com/golang-jwt/jwt"
)

type UserCustomClaims struct {
	Email       string `json:"email"`
	ProfileType string `json:"profileType"`
	jwt.StandardClaims
}

func GenerateToken(email string, profileType string) string {
	mySigningKey := []byte("AllYourBase")
	// usar .env para signKeys
	claims := UserCustomClaims{
		email,
		profileType,
		jwt.StandardClaims{
			Issuer: "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Print(err)
		return ""
	}

	return ss
}

func ParseToken(tokenString string) *UserCustomClaims {
	token, err := jwt.ParseWithClaims(tokenString, &UserCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if claims, ok := token.Claims.(*UserCustomClaims); ok && token.Valid && (err != nil) {
		return claims
	}

	return nil
}
