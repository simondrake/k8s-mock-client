package mockclient

import (
	"context"
	"fmt"
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func TestUpdate(t *testing.T) {
	c := &Client{}

	p := &corev1.Pod{
		Spec: corev1.PodSpec{
			NodeName: "test1",
		},
	}

	c.Update(context.Background(), p)

	fmt.Println(p.Spec.NodeName)
}
