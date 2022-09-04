package post_category

import (
	"fmt"
	"time"

	"github.com/Dias1c/forum/architecture/models"
)

func (c *PostCategoryRepo) GetByPostID(postId int64) ([]*models.PostCategory, error) {
	rows, err := c.db.Query(`
SELECT c.id, c.name, c.created_at FROM posts_categories pc
JOIN categories c ON pc.category_id = c.id
WHERE pc.post_id = ?`, postId)
	if err != nil {
		return nil, fmt.Errorf("c.db.Query: %w", err)
	}

	categories := []*models.PostCategory{}
	for rows.Next() {
		var strCreatedAt string
		category := &models.PostCategory{}
		err = rows.Scan(&category.Id, &category.Name, &strCreatedAt)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}

		timeCreatedAt, err := time.ParseInLocation(models.TimeFormat, strCreatedAt, time.Local)
		if err != nil {
			return nil, fmt.Errorf("time.Parse: %w", err)
		}
		category.CreatedAt = timeCreatedAt
		categories = append(categories, category)
	}
	return categories, nil
}
