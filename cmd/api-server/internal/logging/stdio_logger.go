package logging

import "fmt"

type StdIOLogger struct{}

func (l *StdIOLogger) Info(message string) {
	fmt.Println("[INFO] " + message)
}
