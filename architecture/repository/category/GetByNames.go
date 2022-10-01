package category

import (
	"fmt"
	"strings"
	"time"

	"github.com/Dias1c/forum/architecture/models"
)

func (c *CategoryRepo) GetByNames(names []string) ([]*models.Category, error) {
	if len(names) == 0 {
		return nil, nil
	}

	iNames := make([]interface{}, len(names))
	for i, v := range names {
		iNames[i] = v
	}

	strQuery := fmt.Sprintf(`SELECT id, title, content, user_id, created_at, updated_at FROM posts
WHERE id IN (%v)`, `?`+strings.Repeat(",?", len(iNames)-1))
	fmt.Println(strQuery)
	preQuery := fmt.Sprintf(strQuery, iNames...)
	rows, err := c.db.Query(preQuery)
	if err != nil {
		return nil, fmt.Errorf("db.Query: %w", err)
	}

	cats := []*models.Category{}
	for rows.Next() {
		var strCreatedAt string
		cat := &models.Category{}
		err = rows.Scan(&cat.Id, &cat.Name, &strCreatedAt)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}

		timeCreatedAt, err := time.ParseInLocation(models.TimeFormat, strCreatedAt, time.Local)
		if err != nil {
			return nil, fmt.Errorf("time.Parse: %w", err)
		}
		cat.CreatedAt = timeCreatedAt

		cats = append(cats, cat)
	}
	return cats, nil
}
