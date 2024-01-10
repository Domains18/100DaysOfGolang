package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

var jwtSecretKey = []byte("AutoHotkey")

type jwtclaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func JwtGenerator(username string, email string) (tokenString string, err error) {
	expirationDate := time.Now().Add(1 * time.Hour)
	claims := &jwtclaims{Email: email, Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationDate.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err = token.SignedString(jwtSecretKey)
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtclaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		},
	)
	if err != nil {
		return err
	}
	claims, allIsGood := token.Claims.(*jwtclaims)
	if !allIsGood {
		err = errors.New("failed to parse claims")
		log.Fatal(err)
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		log.Fatal(err)
	}
	return
}
