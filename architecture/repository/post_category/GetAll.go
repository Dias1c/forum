package post_category

import (
	"fmt"
	"time"

	"github.com/Dias1c/forum/architecture/models"
)

func (p *PostCategoryRepo) GetAll(offset, limit int64) ([]*models.PostCategory, error) {
	if limit == 0 {
		limit = -1
	}

	rows, err := p.db.Query(`
SELECT id, name, created_at FROM posts
LIMIT ? OFFSET ? 
	`, limit, offset)

	if err != nil {
		return nil, fmt.Errorf("p.db.Query: %w", err)
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
