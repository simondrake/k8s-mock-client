package mockclient

import (
	"context"
	"reflect"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Get struct {
	Object client.Object
	Error  error

	Capture struct {
		ParameterCapture
	}
}

func (c *Client) Get(ctx context.Context, key client.ObjectKey, obj client.Object) error {
	c.get.Capture.Called++

	if c.get.Object == nil {
		return c.get.Error
	}

	v := reflect.ValueOf(obj).Elem()
	if reflect.ValueOf(c.get.Object).Kind() == reflect.Pointer {
		v.Set(reflect.ValueOf(c.get.Object).Elem())
	} else {
		v.Set(reflect.ValueOf(c.get.Object))
	}

	return c.get.Error
}

type ReaderCapture struct {
	ParameterCapture
}

type List struct {
	ObjectList client.ObjectList
	Error      error

	Capture ReaderCapture
}

func (c *Client) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	c.list.Capture.Called++

	if c.list.ObjectList == nil {
		return c.list.Error
	}

	v := reflect.ValueOf(list).Elem()
	if reflect.ValueOf(c.list.ObjectList).Kind() == reflect.Pointer {
		v.Set(reflect.ValueOf(c.list.ObjectList).Elem())
	} else {
		v.Set(reflect.ValueOf(c.list.ObjectList))
	}

	return c.list.Error
}
