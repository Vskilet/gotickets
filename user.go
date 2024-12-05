package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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

type UserRegister struct {
	FirstName string   `json:"firstname" form:"firstname"`
	LastName  string   `json:"lastname" form:"lastname"`
	Email     string   `json:"mail" form:"mail"`
	Password  Password `json:"pass" form:"pass"`
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

func HandleRegister(ctx *gin.Context) {
	var payload UserRegister
	err := ctx.Bind(&payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	usr := NewUser(payload.FirstName, payload.LastName, payload.Email, payload.Password)
	db.SetUser(usr)
	ctx.JSON(http.StatusOK, usr)
}

func HandleGetUser(ctx *gin.Context) {
	var payload UserRegister
	err := ctx.Bind(&payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	usr, err := db.GetUserByName(payload.LastName)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if usr.Password != payload.Password {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.JSON(200, gin.H{"auth": true})
}

func HandleGetUser(ctx *gin.Context) {
	name := ctx.Query("name")
	usr, er := db.GetUserByName(name)
	if er != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   ErrUserNotFound,
		})
		return
	}
	log.Printf("%v is here", usr.FirstName)
	ctx.JSON(http.StatusOK, usr)
}
