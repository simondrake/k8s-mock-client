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
	Objects []client.Object
}

type StatusUpdate struct {
	Error error

	Capture StatusUpdateCapture
}

func (s *Status) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	s.update.Capture.Called++

	if s.update.Capture.Objects == nil {
		s.update.Capture.Objects = make([]client.Object, 0)
	}

	s.update.Capture.Objects = append(s.update.Capture.Objects, obj)

	return s.update.Error
}

func (s *Status) UpdateCapture() StatusUpdateCapture {
	return s.update.Capture
}

func (s *Status) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
	return nil
}

func (c *Client) GetStatus() *Status {
	return c.status
}
