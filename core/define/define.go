package define

import "github.com/dgrijalva/jwt-go"

type UserClaim struct {
	Id       int
	Identity string
	Username string
	jwt.StandardClaims
}

var JwtKey = "ryangee-cloud-drive"
