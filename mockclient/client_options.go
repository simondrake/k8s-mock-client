package mockclient

type ClientOption func(*Client)

func WithGet(g Get) ClientOption {
	return func(c *Client) {
		c.get = g
	}
}

func WithList(l List) ClientOption {
	return func(c *Client) {
		c.list = l
	}
}

func WithCreate(b Create) ClientOption {
	return func(c *Client) {
		c.create = b
	}
}

func WithUpdate(u Update) ClientOption {
	return func(c *Client) {
		c.update = u
	}
}

func WithStatus(s *Status) ClientOption {
	return func(c *Client) {
		c.status = s
	}
}
