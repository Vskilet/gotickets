package main

import (
	"errors"
	"log"
)

var (
	ErrUserNotFound = errors.New("User not found in database")
	ErrUnauthorized = errors.New("User not authorized or bad password")
)

type DB struct {
	Users map[string]User
}

func NewDB() *DB {
	return &DB{
		Users: make(map[string]User),
	}
}

func (db *DB) SetUser(usr *User) error {
	db.Users[usr.UUID] = *usr
	log.Printf("%v added to DB", usr.FirstName)
	return nil
}

func (db *DB) GetUserByID(id string) (*User, error) {
	u, err := db.Users[id]
	if err {
		return nil, ErrUserNotFound
	}
	return &u, nil
}

func (db *DB) GetUserByName(name string) (*User, error) {
	for _, usr := range db.Users {
		if usr.LastName == name || usr.FirstName == name {
			return &usr, nil
		}
	}
	return nil, ErrUserNotFound
}

func (db *DB) GetUserByMail(email string) (*User, error) {
	for _, usr := range db.Users {
		if usr.Email == email {
			return &usr, nil
		}
	}
	return nil, ErrUserNotFound
}
