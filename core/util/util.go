package util

import (
	"cloud-drive/core/define"
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func GenerateToken(id int, identity, username string) (string, error) {
	userClaim := define.UserClaim{
		Id:       id,
		Identity: identity,
		Username: username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	signedString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return signedString, nil
}
