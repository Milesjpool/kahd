package logging

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() {
		os.Stdout = old
	}()

	f()

	w.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r)
	return buf.String()
}

func TestStdIOLogger(t *testing.T) {

	t.Run("it Info messages to stdout", func(t *testing.T) {
		message := "Hello, world!"
		expectedOutput := "[INFO] Hello, world!\n"

		logger := &StdIOLogger{}
		logOutput := captureStdout(func() {
			logger.Info(message)
		})

		assert.Equal(t, expectedOutput, logOutput)
	})
}
