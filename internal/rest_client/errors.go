package rest

type NotFound struct {
	resource string
}

func (e NotFound) Error() string {
	return "Resource not found: " + e.resource
}
