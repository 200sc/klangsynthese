package supports

// Unsupported is an error type reporting that a filter was not supported
// by the Audio type it was used on
type Unsupported struct {
	filters []string
}

func NewUnsupported(filters []string) Unsupported {
	return Unsupported{filters}
}

func (un Unsupported) Error() string {
	s := "Unsupported filters: "
	for _, f := range un.filters {
		s += f + " "
	}
	return s
}

// Cons combines two unsupported errors into one
func (un Unsupported) Cons(err error) ConsError {
	// At time of writing, Unsupported is the only ConsError
	// Todo: probably use some established library that has done
	// this sort of work combining errors already
	switch err2 := err.(type) {
	case Unsupported:
		return Unsupported{append(un.filters, err2.filters...)}
	}
	return un
}
