package deepAI

const (
	NudityURL = "https://api.deepai.org/api/nsfw-detector"
)

type Client struct {
	Token string
}

func (c *Client) NewClient(token string) {
	c.Token = token
}
