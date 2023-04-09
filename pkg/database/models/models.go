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

type Snippet struct {
	ID        int
	Title     string
	Content   string
	Expires   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID        int
	Name      string
	Email     string
	Password  []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}
