package post_vote

import (
	"fmt"
	"forum/architecture/models"
)

func (p *PostVoteRepo) Create(vote *models.PostVote) (int64, error) {
	strCreatedAt := vote.CreatedAt.Format(models.TimeFormat)
	row := p.db.QueryRow(`
INSERT INTO users (vote, user_id, post_id, created_at) VALUES
(?, ?, ?, ?) RETURNING id`, vote.Vote, vote.UserId, vote.PostId, strCreatedAt)

	err := row.Scan(&vote.Id)
	switch {
	case err == nil:
		return vote.Id, nil
	}
	return -1, fmt.Errorf("row.Scan: %w", err)
}
