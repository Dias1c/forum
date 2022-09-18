package post

import (
	"errors"
	"fmt"

	"github.com/Dias1c/forum/architecture/models"
	rpost "github.com/Dias1c/forum/architecture/repository/post"
)

func (p *PostService) GetWFullPostByID(id, userId int64) (*models.Post, error) {
	post, err := p.repo.GetWFullPostByID(id, userId)
	switch {
	case err == nil:
		return post, nil
	case errors.Is(err, rpost.ErrNotFound):
		return nil, ErrNotFound
	}
	return nil, fmt.Errorf("p.repo.GetWFullPostByID: %w", err)
}
