package mockclient

import "k8s.io/apimachinery/pkg/runtime"

func (c *Client) Scheme() *runtime.Scheme {
	return &runtime.Scheme{}
}
