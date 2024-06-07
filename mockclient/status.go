package mockclient

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Status struct {
	update StatusUpdate
}

func NewStatus() *Status {
	return &Status{}
}

type StatusUpdateCapture struct {
	ParameterCapture
	Object client.Object
}

type StatusUpdate struct {
	Error error

	Capture StatusUpdateCapture
}

func (s *Status) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	s.update.Capture.Called++
	s.update.Capture.Object = obj

	return s.update.Error
}

func (s *Status) GetUpdateCapture() StatusUpdateCapture {
	return s.update.Capture
}

func (s *Status) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
	return nil
}

func (c *Client) GetStatus() *Status {
	return c.status
}
