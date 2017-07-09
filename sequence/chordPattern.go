package sequence

import (
	"time"

	"github.com/200sc/klangsynthese/synth"
)

type ChordPattern struct {
	Pitches [][]synth.Pitch
	Holds   [][]time.Duration
}
