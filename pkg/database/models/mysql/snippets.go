package mysql

import (
	"database/sql"

	"github.com/yousifsabah0/snippetsbox/pkg/database/models"
)

type SnippetModel struct {
	Db *sql.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
