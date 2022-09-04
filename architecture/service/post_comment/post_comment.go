package post_comment

import (
	"github.com/Dias1c/forum/architecture/models"
)

type PostCommentService struct {
	repo models.IPostCommentRepo
}

func NewPostCommentService(repo models.IPostCommentRepo) *PostCommentService {
	return &PostCommentService{repo}
}
