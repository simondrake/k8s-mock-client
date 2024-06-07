package mockclient

import (
	"k8s.io/apimachinery/pkg/api/meta"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ client.Client = &Client{}

type ParameterCapture struct {
	Called int
}

type Client struct {
	get    Get
	list   List
	create Create
	update Update
	status *Status
}

func NewClient(opts ...ClientOption) *Client {
	c := &Client{}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Client) Status() client.StatusWriter {
	return c.status
}

func (c *Client) RESTMapper() meta.RESTMapper {
	return NewRestMapper()
}
