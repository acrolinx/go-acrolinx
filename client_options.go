package acrolinx

type ClientOptionFunc func(*Client) error

func WithAPIToken(token string) ClientOptionFunc {
	return func(c *Client) error {
		c.setToken(token)
		return nil
	}
}
