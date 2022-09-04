package post_category

import "fmt"

func (c *PostCategoryService) DeleteByPostID(postId int64) error {
	err := c.repo.DeleteByPostID(postId)
	switch {
	case err == nil:
	case err != nil:
		return fmt.Errorf("c.repo.DeleteByPostID: %w", err)
	}
	return nil
}
