package mockclient

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type CreateCapture struct {
	ParameterCapture
	Objects []client.Object
}

type Create struct {
	Error error

	Capture CreateCapture
}

func (c *Client) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	c.create.Capture.Called++

	if c.create.Capture.Objects == nil {
		c.create.Capture.Objects = make([]client.Object, 0)
	}

	c.create.Capture.Objects = append(c.create.Capture.Objects, obj)

	return c.create.Error
}

func (c *Client) CreateCapture() CreateCapture {
	return c.create.Capture
}

func (c *Client) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
	return nil
}

type UpdateCapture struct {
	ParameterCapture
	Objects []client.Object
}

type Update struct {
	Error error

	Capture UpdateCapture
}

func (c *Client) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	c.update.Capture.Called++

	if c.update.Capture.Objects == nil {
		c.update.Capture.Objects = make([]client.Object, 0)
	}

	c.update.Capture.Objects = append(c.update.Capture.Objects, obj)

	return c.update.Error
}

func (c *Client) UpdateCapture() UpdateCapture {
	return c.update.Capture
}

func (c *Client) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
	return nil
}

func (c *Client) DeleteAllOf(ctx context.Context, obj client.Object, opts ...client.DeleteAllOfOption) error {
	return nil
}
