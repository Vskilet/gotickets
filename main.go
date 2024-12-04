package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserRegister struct {
	FirstName string   `json:"firstname" form:"firstname"`
	LastName  string   `json:"lastname" form:"lastname"`
	Email     string   `json:"mail" form:"mail"`
	Password  Password `json:"pass" form:"pass"`
}

type Password string

type User struct {
	UUID      string    `json:"uuid"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"mail"`
	Password  Password  `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// This is a reimplementation of a JSON parsing for a Password
func (pwd *Password) UnmarshalJSON(b []byte) error {
	aux := ""
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return err
	}
	h := sha256.New()
	h.Write([]byte(aux))
	*pwd = Password(fmt.Sprintf("%x", h.Sum(nil)))
	return nil
}

func NewUser(firstname, lastname, mail string, pass Password) *User {
	return &User{
		UUID:      uuid.New().String(),
		FirstName: firstname,
		LastName:  lastname,
		Email:     mail,
		Password:  pass,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
func main() {
	router := gin.Default()
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"answer": "ok",
			"code":   http.StatusOK,
		})
	})
	router.POST("/register", HandleRegister)
	router.Run(":8888") // default, listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func HandleRegister(ctx *gin.Context) {
	var payload UserRegister
	err := ctx.Bind(&payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	usr := NewUser(payload.FirstName, payload.LastName, payload.Email, payload.Password)
	ctx.JSON(http.StatusOK, usr)
}
