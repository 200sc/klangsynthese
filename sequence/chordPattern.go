package sequence

import (
	"time"

	"github.com/200sc/klangsynthese/synth"
)

type ChordPattern struct {
	Pitches [][]synth.Pitch
	Holds   [][]time.Duration
}

type HasChords interface {
	GetChordPattern() ChordPattern
	SetChordPattern(ChordPattern)
}

func (cp *ChordPattern) GetChordPattern() ChordPattern {
	return *cp
}

func (cp *ChordPattern) SetChordPattern(cs ChordPattern) {
	*cp = cs
}

// Pitches sets the generator's pitch pattern
func Chords(cp ChordPattern) Option {
	return func(g Generator) {
		if hcp, ok := g.(HasChords); ok {
			hcp.SetChordPattern(cp)
		}
	}
}
