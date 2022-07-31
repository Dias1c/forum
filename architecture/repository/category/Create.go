package category

import (
	"fmt"
	"forum/architecture/models"
	"strings"
)

func (c *CategoryRepo) Create(category *models.Category) (int64, error) {
	row := c.db.QueryRow(`
	INSERT INTO categories (name) VALUES
	(?) RETURNING id`, category.Name)
	err := row.Scan(&category.Id)

	switch {
	case err == nil:
		return category.Id, nil
	case strings.HasPrefix(err.Error(), "UNIQUE constraint failed"):
		return -1, ErrExistName
	case strings.HasPrefix(err.Error(), "CHECK constraint failed"):
		return -1, ErrCheckLengthName
	}
	return -1, fmt.Errorf("row.Scan: %w", err)
}
