package post_category

import (
	"fmt"

	"github.com/Dias1c/forum/architecture/models"
)

func (p *PostCategoryService) GetAll(offset, limit int64) ([]*models.PostCategory, error) {
	categories, err := p.repo.GetAll(offset, limit)
	switch {
	case err == nil:
	case err != nil:
		return nil, fmt.Errorf("p.repo.GetAll: %w", err)
	}
	return categories, nil
}
