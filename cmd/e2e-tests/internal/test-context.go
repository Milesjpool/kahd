package internal

import (
	"fmt"
	"os"
)

type TestContext struct {
	failed bool
}

func (t *TestContext) Run(name string, f func(t *TestContext)) {
	fmt.Printf("Running test: %s\n", name)
	f(t)
}

func (t *TestContext) Fail(format string, args ...interface{}) {
	t.failed = true
	fmt.Printf(format, args...)
}

func (t TestContext) Failed() bool {
	return t.failed
}

func (t *TestContext) Close() {
	if t.Failed() {
		os.Exit(1)
	} else {
		fmt.Println("All tests passed")
		os.Exit(0)
	}
}
