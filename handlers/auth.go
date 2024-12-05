package handlers

import (
	"fmt"
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
		"iss": uuid,
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))

	log.Println(tokenString, err)
	return tokenString, err
}

func (h *Handler) VerifyJWTToken(secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtValue := ctx.GetHeader("Authorization")
		if jwtValue == "" || len(jwtValue) < 7 || jwtValue[:7] != "Bearer " {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		tokenValue := jwtValue[7:]
		checkJWT := &parseMethod{
			secret: secret,
		}
		token, err := jwt.Parse(tokenValue, checkJWT.parser)
		if err != nil {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			log.Println("token is valid", claims["iat"], claims["exp"], claims["iss"])
			ctx.Next()
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

type parseMethod struct {
	secret string
}

func (p *parseMethod) parser(token *jwt.Token) (interface{}, error) {
	// Don't forget to validate the alg is what you expect:
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(p.secret), nil
}
