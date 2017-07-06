package supports

// A ConsError is an error that can be combined with another error
type ConsError interface {
	error
	Cons(e error) ConsError
}
