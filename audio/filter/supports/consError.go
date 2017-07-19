package supports

// A ConsError is an error that can be combined with another error
// Todo: use errors.Wrap instead of this
type ConsError interface {
	error
	Cons(e error) ConsError
}
