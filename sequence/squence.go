package sequence

type Sequence struct{}

type Option func(*Sequence)

func defaultSequence() *Sequence {
	s := &Sequence{}
	return s
}

func New(opts ...Option) *Sequence {
	s := defaultSequence()
	for _, opt := range opts {
		opt(s)
	}
	return s
}
