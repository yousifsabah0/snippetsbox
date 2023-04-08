package mysql

import (
	"database/sql"
	"errors"

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
	snippet := &models.Snippet{}
	query := "SELECT title, content, expires, updated_at FROM snippets WHERE expires > UTC_TIMESTAMP() AND id = ?"

	row := m.Db.QueryRow(query, id)
	if err := row.Scan(&snippet.Title, &snippet.Content, &snippet.Expires, &snippet.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}

		return nil, err
	}

	return snippet, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	snippets := []*models.Snippet{}
	query := "SELECT * FROM snippets WHERE expires > UTC_TIMESTAMP() ORDER BY created_at DESC LIMIT 10"

	rows, err := m.Db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var snippet *models.Snippet
		if err := rows.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Expires, &snippet.CreatedAt, &snippet.UpdatedAt); err != nil {
			return nil, err
		}

		snippets = append(snippets, snippet)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
