package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	client := NewClient("localhost:8080")

	assert.NotNil(t, client)
}
