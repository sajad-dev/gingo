package server

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHttp(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	err := make(chan error, 1)
	go func() {
		close := make(chan struct{})
		erre := Http(8080, nil, close)
		err <- erre
	}()
	select {
	case <-ctx.Done():
		t.Log("Server running succsefuly ;) ")
	case errHttp := <-err:
		assert.NoError(t, errHttp)
	}
}
