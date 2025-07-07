package connection

import (
	"testing"

	"github.com/sajad-dev/gingo/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestConnect (t *testing.T) {
	config.BootConfig("../../../.env")
	assert.NoError(t,Connect())

	if DB == nil {
		t.Fatal("Database connection not set")
	}
}
