package db

import (
	"errors"
	"log"

	"gotickets/models"
)

var (
	ErrUserNotFound = errors.New("User not found in database")
	ErrUnauthorized = errors.New("User not authorized or bad password")
)

type DB struct {
	Users map[string]models.User
}

func NewDB() *DB {
	return &DB{
		Users: make(map[string]models.User),
	}
}

func (db *DB) SetUser(usr *models.User) error {
	db.Users[usr.UUID] = *usr
	log.Printf("%v added to DB", usr.FirstName)
	return nil
}

func (db *DB) GetUserByID(id string) (*models.User, error) {
	u, err := db.Users[id]
	if err {
		return nil, ErrUserNotFound
	}
	return &u, nil
}

func (db *DB) GetUserByName(name string) (*models.User, error) {
	for _, usr := range db.Users {
		if usr.LastName == name || usr.FirstName == name {
			return &usr, nil
		}
	}
	return nil, ErrUserNotFound
}

func (db *DB) GetUserByMail(email string) (*models.User, error) {
	for _, usr := range db.Users {
		if usr.Email == email {
			return &usr, nil
		}
	}
	return nil, ErrUserNotFound
}
