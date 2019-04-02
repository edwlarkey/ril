package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Article struct {
	ID        int `storm:"id,increment"`
	Title     string
	Content   string
	URL       string
	Created   time.Time
	Completed int
}

type User struct {
	ID             int `storm:"id,increment"`
	Name           string
	Email          string `storm:"index"`
	HashedPassword []byte
	Created        time.Time
}
