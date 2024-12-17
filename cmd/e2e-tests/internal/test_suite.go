package internal

type TestSuite interface {
	Run(t *TestContext)
}
