package post

import (
	"fmt"
	"forum/architecture/models"
	"strings"
	"time"

	model "forum/architecture/models"
)

func (p *PostRepo) GetByID(id int64) (*model.Post, error) {
	row := p.db.QueryRow(`
	SELECT id, title, content, user_id, created_at FROM posts
	WHERE id = ?`, id)

	post := &models.Post{}
	strCreatedAt := ""
	err := row.Scan(&post.Id, &post.Title, &post.Content, &post.UserId, &strCreatedAt)
	switch {
	case err == nil:
		timeCreatedAt, err := time.ParseInLocation(models.TimeFormat, strCreatedAt, time.Local)
		if err != nil {
			return nil, fmt.Errorf("time.Parse: %w", err)
		}
		post.CreatedAt = timeCreatedAt
		return post, nil
	case strings.HasPrefix(err.Error(), "sql: no rows in result set"):
		return nil, ErrNotFound
	}
	return nil, fmt.Errorf("row.Scan: %w", err)
}
