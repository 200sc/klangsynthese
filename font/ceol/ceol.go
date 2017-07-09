package ceol

import (
	"time"

	"github.com/200sc/klangsynthese/sequence"
	"github.com/200sc/klangsynthese/synth"
)

// Raw Ceol types, holds all information in ceol file

type Ceol struct {
	Version       int
	Swing         int
	Effect        int
	EffectValue   int
	Bpm           int
	PatternLength int
	BarLength     int
	Instruments   []Instrument
	Patterns      []Pattern
	LoopStart     int
	LoopEnd       int
	Arrangement   [][8]int
}

type Instrument struct {
	Index        int
	IsDrumkit    int
	Palette      int
	LPFCutoff    int
	LPFResonance int
	Volume       int
}

type Pattern struct {
	Key        int
	Scale      int
	Instrument int
	Palette    int
	Notes      []Note
	Filters    []Filter
}

type Note struct {
	PitchIndex int // C4 = 60
	Length     int
	Offset     int
}

type Filter struct {
	Volume       int
	LPFCutoff    int
	LPFResonance int
}

func (c Ceol) ChordPattern() sequence.ChordPattern {
	chp := sequence.ChordPattern{}
	chp.Pitches = make([][]synth.Pitch, c.PatternLength*len(c.Arrangement))
	chp.Holds = make([][]time.Duration, c.PatternLength*len(c.Arrangement))
	for i, m := range c.Arrangement {
		for _, p := range m {
			if p != -1 {
				for _, n := range c.Patterns[p].Notes {
					chp.Pitches[n.Offset+i*c.PatternLength] =
						append(chp.Pitches[n.Offset+i*c.PatternLength], synth.NoteFromIndex(n.PitchIndex))
					chp.Holds[n.Offset+i*c.PatternLength] =
						append(chp.Holds[n.Offset+i*c.PatternLength], DurationFromQuarters(c.Bpm, n.Length))
				}
			}
		}
	}
	return chp
}

func DurationFromQuarters(bpm, quarters int) time.Duration {
	beatTime := time.Duration(60000/bpm) * time.Millisecond
	quarterTime := beatTime / 4
	return quarterTime * time.Duration(quarters)
}
