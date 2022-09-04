package post

import (
	"errors"
	"fmt"

	model "github.com/Dias1c/forum/architecture/models"
	rpost "github.com/Dias1c/forum/architecture/repository/post"
)

func (p *PostService) GetByID(id int64) (*model.Post, error) {
	post, err := p.repo.GetByID(id)
	switch {
	case err == nil:
		return post, nil
	case errors.Is(err, rpost.ErrNotFound):
		return nil, ErrNotFound
	}
	return nil, fmt.Errorf("p.repo.GetByID: %w", err)
}
