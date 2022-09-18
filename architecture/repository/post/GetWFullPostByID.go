package post

import (
	"fmt"
	"strings"
	"time"

	"github.com/Dias1c/forum/architecture/models"
)

func (p *PostRepo) GetWFullPostByID(id, userId int64) (*models.Post, error) {
	row := p.db.QueryRow(`
SELECT 
	p.id as p_id, 
    p.title as p_title, 
    p.content as p_content, 
    p.user_id as p_author_id, 
    p.created_at as p_created_at, 
    p.updated_at as p_updated_at, 
    (SELECT COUNT(vote) FROM posts_votes WHERE post_id = p.id AND vote == 1) as p_vote_up,
    (SELECT COUNT(vote) FROM posts_votes WHERE post_id = p.id AND vote == -1) as p_vote_down,
    u.nickname as p_author_nickname,
    u.email as p_author_email,
    u.created_at as p_author_created_at,
    (SELECT vote FROM posts_votes WHERE post_id = p.id AND user_id = ?) as t_user_vote
FROM posts p
JOIN users u ON u.id = p.user_id
WHERE p.id = ?`, userId, id)

	user := &models.User{}
	post := &models.Post{}
	var postStrCreatedAt, postStrUpdatedAt, userStrCreatedAt string
	err := row.Scan(&post.Id, &post.Title, &post.Content, &post.UserId, &postStrCreatedAt, &postStrUpdatedAt, &user.Nickname, &user.Email, &userStrCreatedAt, &post.WFVoteUp, &post.WFVoteDown, &post.WFUserVote)
	switch {
	case err == nil:
	case strings.HasPrefix(err.Error(), "sql: no rows in result set"):
		return nil, ErrNotFound
	case err != nil:
		return nil, fmt.Errorf("row.Scan: %w", err)
	}
	post.WFUser = user

	timeCreatedAt, err := time.ParseInLocation(models.TimeFormat, postStrCreatedAt, time.Local)
	if err != nil {
		return nil, fmt.Errorf("time.Parse: %w", err)
	}
	post.CreatedAt = timeCreatedAt

	timeUpdatedAt, err := time.ParseInLocation(models.TimeFormat, postStrUpdatedAt, time.Local)
	if err != nil {
		return nil, fmt.Errorf("time.Parse: %w", err)
	}
	post.UpdatedAt = timeUpdatedAt

	timeUsrCreatedAt, err := time.ParseInLocation(models.TimeFormat, userStrCreatedAt, time.Local)
	if err != nil {
		return nil, fmt.Errorf("time.Parse: %w", err)
	}
	user.CreatedAt = timeUsrCreatedAt

	return post, nil
}
