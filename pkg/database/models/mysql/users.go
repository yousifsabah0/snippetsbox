package mysql

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/yousifsabah0/snippetsbox/pkg/database/models"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	Db *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	_, err = m.Db.Exec(query, name, email, string(hash))
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.Is(err, mysqlErr) {
			if mysqlErr.Number == 1062 && strings.Contains(mysqlErr.Message, "users_uc_email") {
				return err
			}
		}

		return err
	}

	return nil
}

func (m *UserModel) Authenticate(email, pass string) (int, error) {
	var id int
	var password []byte

	query := "SELECT id, password FROM users WHERE email = ?"
	row := m.Db.QueryRow(query, email)

	if err := row.Scan(&id, &password); err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			return 0, models.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	if err := bcrypt.CompareHashAndPassword(password, []byte(pass)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, models.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	return id, nil
}

func (m *UserModel) Get(id string) (*models.User, error) {
	return nil, nil
}
