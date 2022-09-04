package models

import "fmt"

func (p *Post) ValidateTitle() error {
	if lng := len(p.Title); lng < 1 || 100 < lng {
		return fmt.Errorf("title: invalid lenght (%d)", lng)
	}
	return nil
}

func (p *Post) ValidateContent() error {
	if lng := len(p.Content); lng < 1 {
		return fmt.Errorf("content: invalid lenght (%d)", lng)
	}
	return nil
}
