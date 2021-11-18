package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Mr-Malomz/skintech_be/dtos"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		if header != "" {
			splitHeader := strings.Split(header, " ")
			if len(splitHeader) == 2 {
				reqToken := splitHeader[1]

				tkn, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
					//validate the algorithm
					if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
					}
					return []byte("secret"), nil
				})

				if err != nil {
					c.JSON(http.StatusForbidden, dtos.Response{Status: http.StatusForbidden, Message: "invalid authorization token"})
					// c.AbortWithStatus(http.StatusForbidden)
				}

				if !tkn.Valid {
					// c.Next()
					c.JSON(http.StatusForbidden, dtos.Response{Status: http.StatusForbidden, Message: "invalid authorization token"})
				} 
			}
		} else {
			c.JSON(http.StatusForbidden, dtos.Response{Status: http.StatusForbidden, Message: "An authorization header is required"})
			c.AbortWithStatus(http.StatusForbidden)
		}

	}
}
