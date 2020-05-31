package ukco

// Client is the struct where all http calls are made
type Client struct {
	APIKey string
}

// Get performs a http get against the Companies House API
func (c *Client) Get(path string) {

}
