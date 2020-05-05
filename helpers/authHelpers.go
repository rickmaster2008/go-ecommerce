package helpers

import (
	"newproject/env"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//Claims defines bse claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// ObtainToken generates new token for user
func ObtainToken(username string) (string, error) {
	appKey := env.GetVariable("APP_KEY")
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(appKey))
}

//VerifyToken if is valid
func VerifyToken(tokenString string) (Claims, bool, error) {
	claims := Claims{}
	appKey := env.GetVariable("APP_KEY")
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(tkn *jwt.Token) (interface{}, error) {
		return []byte(appKey), nil
	})

	if err != nil {
		return claims, false, err
	}
	return claims, token.Valid, nil
}
