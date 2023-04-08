package mysql

import (
	"database/sql"

	"github.com/yousifsabah0/snippetsbox/pkg/database/models"
)

type SnippetModel struct {
	Db *sql.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	query := "INSERT INTO snippets (title, content, expires) VALUES (?, ?, DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))"
	result, err := m.Db.Exec(query, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	var snippet *models.Snippet
	query := "SELECT title, content, expires, updated_at FROM snippets WHERE id = ?"

	row := m.Db.QueryRow(query, id)
	if err := row.Scan(&snippet.Title, &snippet.Content, &snippet.Expires, &snippet.UpdatedAt); err != nil {
		return nil, err
	}

	return snippet, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	var snippets []*models.Snippet
	query := "SELECT * FROM snippets LIMIT 10"

	rows, err := m.Db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var snippet *models.Snippet
		if err := rows.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Expires, &snippet.CreatedAt, &snippet.UpdatedAt); err != nil {
			return nil, err
		}

		snippets = append(snippets, snippet)
	}

	return snippets, nil
}
