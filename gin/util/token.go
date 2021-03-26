package token

import (
	"main/conf"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	jwt.StandardClaims

	Uid   uint `json:"uid"`
	Admin bool `json:"admin"`
}

func CreateToken(SecretKey []byte, issuer string, uid uint, isAdmin bool) (string, error) {
	claims := &UserClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(conf.ExpiresAt).Unix()),
			Issuer:    issuer,
		},
		uid,
		isAdmin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}

func ParseToken(tokenString string, SecretKey []byte) (bool, uint, string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return false, 0, "Invalid token"
	}
	var flag bool
	var uid uint
	var msg string
	if token.Valid {
		flag = true
		claims := token.Claims
		msg = "Pass"
		uid = uint(claims.(jwt.MapClaims)["uid"].(float64))
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		flag = false
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			msg = "That's not even a token"
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			msg = "Timing is everything"
		} else {
			msg = "Couldn't handle this token"
		}
	} else {
		flag = false
		msg = "Couldn't handle this token"
	}
	return flag, uid, msg
}
