package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func BasicAuth(ctx *gin.Context) {
	user, pass, ok := ctx.Request.BasicAuth()
	if !ok || user != "admin" || pass != "admin" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	ctx.Next()
}

func (h *Handler) newJWTToken(uuid string, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":      uuid,
		"notBefore": time.Now().Unix(),
		"notAfter":  time.Now().Add(time.Minute * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secret)

	log.Println(tokenString, err)
	return tokenString, err
}
