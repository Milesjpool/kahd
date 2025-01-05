package logging

import "fmt"

type MockLogger struct {
	InfoCalls  [][]string
	InfoLogs   []string
	ErrorCalls [][]string
	ErrorLogs  []string
}

func stringifyArgs(args ...any) []string {
	stringArgs := make([]string, len(args))
	for i, arg := range args {
		stringArgs[i] = fmt.Sprint(arg)
	}
	return stringArgs
}

func (m *MockLogger) Info(message string, args ...any) {
	m.InfoCalls = append(m.InfoCalls, append([]string{message}, stringifyArgs(args...)...))
	m.InfoLogs = append(m.InfoLogs, fmt.Sprintf(message, args...))
}

func (m *MockLogger) Error(message string, args ...any) {
	m.ErrorCalls = append(m.ErrorCalls, append([]string{message}, stringifyArgs(args...)...))
	m.ErrorLogs = append(m.ErrorLogs, fmt.Sprintf(message, args...))
}
