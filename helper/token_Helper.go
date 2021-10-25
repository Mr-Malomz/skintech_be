package helper

import (
	"log"
	"time"

	"github.com/Mr-Malomz/skintech_be/config"
	jwt "github.com/dgrijalva/jwt-go"
)

type Payload struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(id string, firstname string, lastname string, email string) (token string, refreshToken string, err error) {
	claims := &Payload{
		ID:        id,
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(15)).Unix(),
		},
	}

	refreshClaims := &Payload{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	//creating the token
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.EnvJWT()))

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(config.EnvJWT()))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}
