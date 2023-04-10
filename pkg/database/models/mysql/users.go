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

func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (m *UserModel) Get(id string) (*models.User, error) {
	return nil, nil
}
