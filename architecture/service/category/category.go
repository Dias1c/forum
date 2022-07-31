package category

import (
	"forum/architecture/models"
)

type CategoryService struct {
	repo models.ICategoryRepo
}

func NewCategoryService(repo models.ICategoryRepo) *CategoryService {
	return &CategoryService{repo}
}
