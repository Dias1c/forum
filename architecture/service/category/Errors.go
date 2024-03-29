package category

import (
	"errors"
	"fmt"

	"github.com/Dias1c/forum/architecture/models"
)

var (
	ErrCategoryLimitForPost = fmt.Errorf("category limit for post greater than %v", models.MaxCategoryLimitForPost)
	ErrExistName            = errors.New("category with this name exists")
	ErrCheckLengthName      = errors.New("category name length is wrong")
	ErrNotFound             = errors.New("category not found")
)
