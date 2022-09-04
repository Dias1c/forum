package post_category

import (
	"fmt"
	"forum/architecture/models"
)

func (c *PostCategoryService) GetByPostID(postId int64) ([]*models.PostCategory, error) {
	categories, err := c.repo.GetByPostID(postId)
	switch {
	case err == nil:
		return categories, nil
	}
	return nil, fmt.Errorf("c.repo.GetByPostID: %w", err)
}
