package book

// CreateBookDTO represents needed fields to create book.
// It is needed for Service interface and is linking chain between business logic and controller layers.
type CreateBookDTO struct {
	Name       string `json:"name"`
	Year       int    `json:"year"`
	AuthorUUID string `json:"author_uuid"`
}

// UpdateBookDTO represents needed fields to update book.
// It is needed for Service interface.
type UpdateBookDTO struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	Year       int    `json:"year"`
	AuthorUUID string `json:"author_uuid"`
	Busy       bool   `json:"busy"`
	Owner      string `json:"owner"`
}
