package sequence

// This is notes / pseudo-code / not useable yet

// A Sequence does not care if it loops because that is audio/Encoding's job
// A Sequence does not care how long it should play each sample it is given
// because that is the job of the individual samples
// A Sequence does care how much time it should wait between samples
// A Sequence does care if that time is variable (swing rythm)
// A Sequence satisfies Audio
type Sequence struct{}

// A Generator stores settings to create a sequence
type Generator struct{}

type Option func(*Generator)

func defaultGen() *Generator {
	g := &Generator{}
	return g
}

func New(opts ...Option) *Generator {
	g := defaultGen()
	for _, opt := range opts {
		opt(g)
	}
	return g
}

func (g *Generator) Generate() *Sequence {
	//...
	return &Sequence{}
}
