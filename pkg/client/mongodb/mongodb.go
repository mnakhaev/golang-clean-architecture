package mongodb

type Client struct {
}

// NewClient returns some DB client.
func NewClient() *Client {
	return &Client{}
}
