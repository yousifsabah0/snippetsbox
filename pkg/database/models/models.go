package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord = errors.New("models: no matching record found")
)

type Snippet struct {
	ID        int
	Title     string
	Content   string
	Expires   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
