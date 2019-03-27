package aci

// Client is the client for all ACI operations.
// Create one of these per session, and always use NewClient to create one
type Client struct {
	token string
}

// NewClient creates a new ACI client with the given token
func NewClient(token string) *Client {
	return &Client{
		token: token,
	}
}
