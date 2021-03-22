package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtCustomerClaim struct {
	jwt.StandardClaims

	Uid   uint `json:"uid"`
	Admin bool `json:"admin"`
}

func CreateToken(SecretKey []byte, issuer string, Uid uint, isAdmin bool) (string, error) {
	claims := &jwtCustomerClaim{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 72).Unix()),
			Issuer:    issuer,
		},
		Uid,
		isAdmin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}

func ParseToken(tokenString string, SecretKey []byte) (bool, string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	var flag bool
	var msg string
	if token.Valid {
		flag = true
		msg = "Token Pass"
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			msg = "That's not even a token"
			flag = false
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			msg = "Timing is everything"
			flag = false
		} else {
			msg = "Couldn't handle this token"
			flag = false
		}
	}
	return flag, msg
}
