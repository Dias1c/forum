package models

import "fmt"

func (c *PostComment) ValidateContent() error {
	if lng := len(c.Content); lng < 1 {
		return fmt.Errorf("content: invalid lenght (%d)", lng)
	}
	return nil
}
