package models

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Password string

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

func (pwd *Password) String() string {
	return "********"
}

type User struct {
	UUID      string    `json:"uuid"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"mail"`
	Password  Password  `json:"pass"`
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
