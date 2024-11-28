package internal

import (
	"fmt"
	"os"
)

type TestContext struct {
	fatal  bool
	failed bool
}

func (t *TestContext) Init(f func(t *TestContext)) {
	fmt.Printf("Starting E2E tests\n")
	f(t)
}

func (t *TestContext) Run(name string, f func(t *TestContext)) {
	if t.fatal {
		return
	}
	fmt.Printf("⏺ %s\n", name)
	f(t)
}

func (t *TestContext) Fail(format string, args ...interface{}) {
	t.failed = true
	fmt.Printf("  × "+format+"\n", args...)
}

func (t *TestContext) Fatal(format string, args ...interface{}) {
	t.fatal = true
	fmt.Printf("△ Fatal: "+format+"\n", args...)
}

func (t TestContext) Failed() bool {
	return t.fatal || t.failed
}

func (t *TestContext) Close() {
	if t.Failed() {
		os.Exit(1)
	} else {
		fmt.Println("★ All tests passed")
		os.Exit(0)
	}
}
