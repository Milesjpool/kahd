package logging

type Logger interface {
	Info(message string, args ...any)
	Error(message string, args ...any)
}
