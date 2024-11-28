package assertions

import (
	"e2e-tests/internal"
	"fmt"
)

func Equals(t *internal.TestContext, expected, actual interface{}, msgArgs ...interface{}) {
	if expected != actual {
		if len(msgArgs) == 0 {
			t.Fail("expected: %v, got: %v", expected, actual)
		} else {
			t.Fail(fmt.Sprint(msgArgs[0]), msgArgs[1:]...)
		}
	}
}

func NoErr(t *internal.TestContext, err error, msg string, msgArgs ...interface{}) {
	if err != nil {
		t.Fail(msg, msgArgs...)
	}
}
