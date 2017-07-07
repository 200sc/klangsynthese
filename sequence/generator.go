package sequence

// A Generator stores settings to create a sequence
type Generator interface {
	Generate() *Sequence
}

// Option types are inserted into Constructors to create generators
type Option func(Generator)
