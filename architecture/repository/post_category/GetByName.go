package post_category

import (
	"fmt"
	"forum/architecture/models"
	"strings"
	"time"
)

func (c *PostCategoryRepo) GetByName(name string) (*models.PostCategory, error) {
	row := c.db.QueryRow(`
SELECT id, name, created_at FROM categories
WHERE name = ?`, name)

	category := &models.PostCategory{}
	var strCreatedAt string
	err := row.Scan(&category.Id, &category.Name, &strCreatedAt)
	switch {
	case err == nil:
		timeCreatedAt, err := time.ParseInLocation(models.TimeFormat, strCreatedAt, time.Local)
		if err != nil {
			return nil, fmt.Errorf("time.Parse: %w", err)
		}
		category.CreatedAt = timeCreatedAt
		return category, nil
	case strings.HasPrefix(err.Error(), "sql: no rows in result set"):
		return nil, ErrNotFound
	}
	return nil, fmt.Errorf("row.Scan: %w", err)
}