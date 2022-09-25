package post

import (
	"fmt"

	"github.com/Dias1c/forum/architecture/models"
)

func (p *PostService) GetByIds(ids []int64) ([]*models.Post, error) {
	posts, err := p.repo.GetByIds(ids)
	switch {
	case err == nil:
	case err != nil:
		return nil, fmt.Errorf("GetByIds: %w", err)
	}
	return posts, nil
}
