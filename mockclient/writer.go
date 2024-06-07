package mockclient

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (c *Client) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	return nil
}

func (c *Client) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
	return nil
}

type UpdateCapture struct {
	ParameterCapture
	Object client.Object
}

type Update struct {
	Error error

	Capture UpdateCapture
}

func (c *Client) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	c.update.Capture.Called++
	c.update.Capture.Object = obj

	return c.update.Error
}

func (c *Client) GetUpdateCapture() UpdateCapture {
	return c.update.Capture
}

func (c *Client) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
	return nil
}

func (c *Client) DeleteAllOf(ctx context.Context, obj client.Object, opts ...client.DeleteAllOfOption) error {
	return nil
}
