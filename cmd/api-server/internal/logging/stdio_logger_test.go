package logging

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func captureStdIO(f func()) (string, string) {
	oldOut := os.Stdout
	oldErr := os.Stderr

	rOut, wOut, _ := os.Pipe()
	rErr, wErr, _ := os.Pipe()
	os.Stdout = wOut
	os.Stderr = wErr
	defer func() {
		os.Stdout = oldOut
		os.Stderr = oldErr
	}()

	f()

	wOut.Close()
	wErr.Close()

	var bufOut bytes.Buffer
	bufOut.ReadFrom(rOut)
	var bufErr bytes.Buffer
	bufErr.ReadFrom(rErr)

	return bufOut.String(), bufErr.String()
}

func TestStdIOLogger(t *testing.T) {

	t.Run("it Info messages to stdout", func(t *testing.T) {
		message := "Hello, world!"
		expectedOutput := "[INFO] Hello, world!\n"

		logger := &StdIOLogger{}
		logOutput, errOutput := captureStdIO(func() {
			logger.Info(message)
		})

		assert.Equal(t, expectedOutput, logOutput)
		assert.Equal(t, "", errOutput)
	})

	t.Run("it Error messages to stderr", func(t *testing.T) {
		message := "Hello, world!"
		expectedOutput := "[ERROR] Hello, world!\n"

		logger := &StdIOLogger{}
		logOutput, errOutput := captureStdIO(func() {
			logger.Error(message)
		})

		assert.Equal(t, "", logOutput)
		assert.Equal(t, expectedOutput, errOutput)
	})

	t.Run("it Info messages with args", func(t *testing.T) {
		message := "Hello, %s!"
		args := []any{"world"}
		expectedOutput := "[INFO] Hello, world!\n"

		logger := &StdIOLogger{}
		logOutput, errOutput := captureStdIO(func() {
			logger.Info(message, args...)
		})

		assert.Equal(t, expectedOutput, logOutput)
		assert.Equal(t, "", errOutput)
	})
}
