package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	jwt.StandardClaims

	Uid   uint `json:"uid"`
	Admin bool `json:"admin"`
}

func CreateToken(SecretKey []byte, issuer string, ExpiresAt time.Duration, uid uint, isAdmin bool) (string, error) {
	claims := &UserClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(ExpiresAt).Unix()),
			Issuer:    issuer,
		},
		uid,
		isAdmin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}
func main() {
	//和后端的一样
	var SecretKey = []byte("SecretKey")
	var Issuer = []byte("Issuer")
	var ExpiresAt = time.Hour * 24
	uid, _ := strconv.ParseUint(os.Args[1], 10, 0)
	isAdmin, _ := strconv.ParseBool(os.Args[2])
	token, err := CreateToken(SecretKey, string(Issuer), ExpiresAt, uint(uid), isAdmin)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(token)
	}
}
