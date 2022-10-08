package helper

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        interface{}
	jwt.StandardClaims
}


func GenerateAllTokens(email string, firstName string, lastName string, uid interface{}) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email:      email,
		First_name: firstName,
		Last_name:  lastName,
		Uid:        uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("secret"))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte("secret"))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}


func VerifyToken(signedToken string) (claims *SignedDetails, msg string){
	token, err := jwt.ParseWithClaims(
        signedToken,
        &SignedDetails{},
        func(token *jwt.Token) (interface{}, error) {
            return []byte("secret"), nil
        },
    )

    if err != nil {
        msg = err.Error()
        return
    }

	claims, ok := token.Claims.(*SignedDetails)
    if !ok {
        msg = fmt.Sprintf("the token is invalid")
        msg = err.Error()
        return
    }

    if claims.ExpiresAt < time.Now().Local().Unix() {
        msg = fmt.Sprintf("token is expired")
        msg = err.Error()
        return
    }

    return claims, msg
}