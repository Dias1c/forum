package post_comment_vote

import "github.com/Dias1c/forum/architecture/models"

type PostCommentVoteService struct {
	repo models.IPostCommentVoteRepo
}

func NewPostCommentVoteService(postCommentVote models.IPostCommentVoteRepo) *PostCommentVoteService {
	return &PostCommentVoteService{postCommentVote}
}
