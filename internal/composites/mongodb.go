package composites

import "ca-library-app/pkg/client/mongodb"

// MongoDBComposite is needed in order to avoid passing DB objects into book/author composites.
type MongoDBComposite struct {
	db *mongodb.Client
}

// NewMongoDBComposite is just fake method, realization is not important here.
// Ideally it should accept ctx and config parameters (host, port, username,...) .
func NewMongoDBComposite() (*MongoDBComposite, error) {
	client := mongodb.NewClient()
	return &MongoDBComposite{db: client}, nil
}
