package post_category

import (
	"github.com/Dias1c/forum/architecture/models"
)

type PostCategoryService struct {
	repo models.IPostCategoryRepo
}

func NewPostCategoryService(repo models.IPostCategoryRepo) *PostCategoryService {
	return &PostCategoryService{repo}
}
