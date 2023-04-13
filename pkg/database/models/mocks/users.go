package mocks

import (
	"time"

	"github.com/yousifsabah0/snippetsbox/pkg/database/models"
)

var userMock = &models.User{
	ID:        1,
	Name:      "John Doe",
	Email:     "john@doe.com",
	Password:  []byte("john"),
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

type UserModel struct{}

func (u *UserModel) Insert(name, email, password string) error {
	switch email {
	case "john2@doe.com":
		return nil
	case "duplicated@email.com":
		return models.ErrDuplicateEmail
	default:
		return models.ErrInvalidCredentials
	}
}

func (u *UserModel) Authenticate(email, password string) (int, error) {
	switch email {
	case "john@doe.com":
		return 1, nil
	default:
		return 0, models.ErrInvalidCredentials
	}
}

func (u *UserModel) Get(id int) (*models.User, error) {
	switch id {
	case 1:
		return userMock, nil
	default:
		return nil, models.ErrNoRecord
	}
}
