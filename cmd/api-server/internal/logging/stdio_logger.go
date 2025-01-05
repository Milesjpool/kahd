package logging

import (
	"fmt"
	"os"
)

type StdIOLogger struct{}

func (l *StdIOLogger) Info(message string, args ...any) {
	fmt.Fprintf(os.Stdout, "[INFO] "+message+"\n", args...)
}

func (l *StdIOLogger) Error(message string, args ...any) {
	fmt.Fprintf(os.Stderr, "[ERROR] "+message+"\n", args...)
}
