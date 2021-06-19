package util

import (
	"XWS-Nistagram-2021/backend-nistagram/authenticationService/model"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

const secret = "nistagram_secret"

type TokenClaims struct {
	UserId string `json:"userId"`
	Role model.UserRole `json:"role"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func CreateJWT(userId string, role *model.UserRole, username string) (string, error) {
	tokenClaims := TokenClaims{UserId: userId, Role: *role, Username: username, StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 50).Unix(),
		IssuedAt: time.Now().Unix(),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func GetJWT(r http.Header) string {
	bearToken := r.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
	return strArr[1]
	}
	return ""
}

func GetUserIdFromToken(r *http.Request) (string){
	tokenString := GetJWT(r.Header)

	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims.UserId
	} else {
		fmt.Println(err)
		return "";
	}
}