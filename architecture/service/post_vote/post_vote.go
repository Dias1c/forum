package post_vote

import (
	"github.com/Dias1c/forum/architecture/models"
)

type PostVoteService struct {
	repo models.IPostVoteRepo
}

func NewPostVoteService(postVote models.IPostVoteRepo) *PostVoteService {
	return &PostVoteService{postVote}
}
