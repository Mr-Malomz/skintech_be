package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Mr-Malomz/skintech_be/config"
	"github.com/Mr-Malomz/skintech_be/dtos"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		splitHeader := strings.Split(header, "Bearer ")
		reqToken := splitHeader[1]

		if header == "" {
			log.Fatal()
			c.JSON(http.StatusForbidden, dtos.Response{Status: http.StatusInternalServerError, Message: "An authorization header is required"})
			return
		}

		tkn, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
			//validate the algorithm
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(config.EnvJWT()), nil
		})

		if err != nil {
			log.Fatal()
			c.JSON(http.StatusForbidden, dtos.Response{Status: http.StatusInternalServerError, Message: err.Error()})
			return
		}

		if !tkn.Valid {
			log.Fatal()
			c.JSON(http.StatusForbidden, dtos.Response{Status: http.StatusInternalServerError, Message: "nvalid authorization token"})
			return
		}

	}
}
