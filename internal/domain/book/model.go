package book

import (
	"fmt"

	"ca-library-app/internal/domain/author"
)

type Book struct {
	UUID   string        `json:"uuid,omitempty"`
	Name   string        `json:"name,omitempty"`
	Year   int           `json:"year,omitempty"`
	Author author.Author `json:"author,omitempty"`
	Busy   bool          `json:"busy,omitempty"`
	Owner  string        `json:"owner,omitempty"`
}

// Take represents business logic.
func (b *Book) Take(owner string) error {
	if b.Busy {
		return fmt.Errorf("book %s is busy", b.Name)
	}
	b.Owner = owner
	b.Busy = true
	return nil
}
