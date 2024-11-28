package assertions

import "github.com/milesjpool/kahd/cmd/e2e-tests/internal"

func Equals(t *internal.TestContext, expected, actual interface{}, msg string, msgArgs ...interface{}) {
	if expected != actual {
		t.Fail(msg, msgArgs...)
	}
}

func NoErr(t *internal.TestContext, err error, msg string, msgArgs ...interface{}) {
	if err != nil {
		t.Fail(msg, msgArgs...)
	}
}
