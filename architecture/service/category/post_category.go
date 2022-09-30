package category

import (
	"github.com/Dias1c/forum/architecture/models"
)

type CategoryService struct {
	repo models.ICategoryRepo
}

func NewPostCategoryService(repo models.ICategoryRepo) *CategoryService {
	return &CategoryService{repo}
}
