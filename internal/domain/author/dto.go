package author

// CreateAuthorDTO represents needed fields to create Author.
// It is needed for Service interface and is linking chain between business logic and controller layers.
type CreateAuthorDTO struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// UpdateAuthorDTO represents needed fields to update Author.
// It is needed for Service interface.
type UpdateAuthorDTO struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
