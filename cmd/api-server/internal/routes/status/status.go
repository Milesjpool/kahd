package status

type Context struct {
	Checks map[string]func() bool
}
