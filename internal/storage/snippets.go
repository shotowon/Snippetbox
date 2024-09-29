package storage

import (
	"snippetbox/internal/models"
)

func (s *Storage) InsertSnippet(
	title string,
	content string,
	expires int,
) (inserted int, err error) {
	return
}

func (s *Storage) GetSnippet(id int) (*models.Snippet, error) {
	return nil, nil
}

func (s *Storage) LatestSnippets() ([]*models.Snippet, error) {
	return nil, nil
}
