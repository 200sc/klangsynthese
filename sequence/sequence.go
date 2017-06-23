package sequence

import (
	"time"

	"github.com/200sc/klangsynthese/audio"
)

// This is notes / pseudo-code / not useable yet

// A Sequence does not care if it loops because that is audio/Encoding's job
// A Sequence does not care how long it should play each sample it is given
// because that is the job of the individual samples
// A Sequence does care how much time it should wait between samples
// A Sequence does care if that time is variable (swing rhythm)
// A Sequence satisfies Audio
type Sequence struct {
	// Sequences play patterns of audio
	// everything at Pattern[0] will be simultaneously Play()ed at
	// Sequence.Play()
	Pattern [][]audio.Audio
	// Every tick, the next index in Pattern will be played by a Sequence
	// until the pattern is over.
	Ticker time.Ticker
}

// A Generator stores settings to create a sequence
type Generator interface {
	Generate() *Sequence
}

type Option func(Generator)

func defaultGen() Generator {
	return nil
}

func New(opts ...Option) Generator {
	g := defaultGen()
	for _, opt := range opts {
		opt(g)
	}
	return g
}

func Sub(sub Generator) Option {
	return func(parent Generator) {
		// Set sub to be the next portion of
		// parent's pattern
	}
}
