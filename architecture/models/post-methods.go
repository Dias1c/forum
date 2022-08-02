package models

import "fmt"

func (p *Post) ValidateTitle() error {
	if lng := len(p.Title); lng < 1 || 100 < lng {
		return fmt.Errorf("title: invalid lenght (%d)", lng)
	}
	return nil
}
