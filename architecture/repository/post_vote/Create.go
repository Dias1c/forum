package post_vote

import (
	"fmt"
	"forum/architecture/models"
	"strings"
)

func (p *PostVoteRepo) Create(vote *models.PostVote) (int64, error) {
	strCreatedAt := vote.CreatedAt.Format(models.TimeFormat)
	row := p.db.QueryRow(`
INSERT INTO posts_votes (vote, user_id, post_id, created_at, updated_at) VALUES
(?, ?, ?, ?, ?) RETURNING id`, vote.Vote, vote.UserId, vote.PostId, strCreatedAt, strCreatedAt)

	err := row.Scan(&vote.Id)
	switch {
	case err == nil:
	case err != nil:
		switch {
		case strings.HasPrefix(err.Error(), "UNIQUE constraint failed"):
			return -1, ErrVoteExists
		}
		return -1, fmt.Errorf("row.Scan: %w", err)
	}
	return vote.Id, nil
}
