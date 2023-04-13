package mocks

import (
	"time"

	"github.com/yousifsabah0/snippetsbox/pkg/database/models"
)

var mockSnippet = &models.Snippet{
	ID:        1,
	Title:     "Mocked Title",
	Content:   "Mocked Content",
	Expires:   time.Now(),
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

type SnippetModel struct{}

func (s *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 2, nil
}

func (s *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (s *SnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippet}, nil
}
