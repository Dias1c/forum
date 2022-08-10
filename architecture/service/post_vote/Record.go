package post_vote

import (
	"errors"
	"fmt"
	"forum/architecture/models"
	"forum/architecture/repository/post_vote"
	"time"
)

func (p *PostVoteService) Record(vote *models.PostVote) error {
	vote.CreatedAt = time.Now()
	_, err := p.repo.Create(vote)
	switch {
	case err == nil:
		return nil
	case errors.Is(err, post_vote.ErrVoteExists):
	case err != nil:
		return fmt.Errorf("p.repo.Create: %w", err)
	}
	// return err
	vote.UpdatedAt = time.Now()
	err = p.repo.Update(vote)
	switch {
	case err == nil:
	case err != nil:
		return fmt.Errorf("p.repo.Update: %w", err)
	}
	return nil
}
