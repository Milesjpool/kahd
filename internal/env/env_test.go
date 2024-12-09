package env

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOrDefault(t *testing.T) {
	t.Run("existing key", func(t *testing.T) {
		t.Setenv("EXISTING_KEY", "myValue")

		result := GetOrDefault("EXISTING_KEY", "default1")
		assert.Equal(t, "myValue", result)
	})

	t.Run("non-existing key", func(t *testing.T) {
		result := GetOrDefault("NON_EXISTING_KEY", "myDefault")
		assert.Equal(t, "myDefault", result)
	})
}

func TestGet(t *testing.T) {
	t.Run("existing key", func(t *testing.T) {
		t.Setenv("EXISTING_KEY", "myValue")

		result, err := Get("EXISTING_KEY")
		assert.Equal(t, "myValue", result)
		assert.Nil(t, err)
	})

	t.Run("non-existing key", func(t *testing.T) {
		_, err := Get("NON_EXISTING_KEY")
		assert.Error(t, err)
	})
}
