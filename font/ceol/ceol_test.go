package ceol

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

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
	SongLength    int
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

func TestReadCeol(t *testing.T) {
	f, err := os.Open("test.ceol")
	assert.Nil(t, err)
	b, err := ioutil.ReadAll(f)
	assert.Nil(t, err)
	s := string(b)
	in := strings.Split(s, ",")
	ints := make([]int, len(in))
	for i, s := range in {
		ints[i], err = strconv.Atoi(s)
		assert.Nil(t, err)
	}
	c := Ceol{}
	i := 0
	c.Version = ints[i]
	i++
	c.Swing = ints[i]
	i++
	c.Effect = ints[i]
	i++
	c.EffectValue = ints[i]
	i++
	c.Bpm = ints[i]
	i++
	c.PatternLength = ints[i]
	i++
	c.BarLength = ints[i]
	i++
	nInstruments := ints[i]
	i++
	c.Instruments = make([]Instrument, nInstruments)
	for j := 0; j < nInstruments; j++ {
		c.Instruments[j].Index = ints[i]
		i++
		c.Instruments[j].IsDrumkit = ints[i]
		i++
		c.Instruments[j].Palette = ints[i]
		i++
		c.Instruments[j].LPFCutoff = ints[i]
		i++
		c.Instruments[j].LPFResonance = ints[i]
		i++
		c.Instruments[j].Volume = ints[i]
		i++
	}
	nPatterns := ints[i]
	i++
	c.Patterns = make([]Pattern, nPatterns)
	for j := 0; j < nPatterns; j++ {
		c.Patterns[j].Key = ints[i]
		i++
		c.Patterns[j].Scale = ints[i]
		i++
		c.Patterns[j].Instrument = ints[i]
		i++
		c.Patterns[j].Palette = ints[i]
		i++
		nNotes := ints[i]
		i++
		c.Patterns[j].Notes = make([]Note, nNotes)
		for k := 0; k < nNotes; k++ {
			c.Patterns[j].Notes[k].PitchIndex = ints[i]
			i++
			c.Patterns[j].Notes[k].Length = ints[i]
			i++
			c.Patterns[j].Notes[k].Offset = ints[i]
			i++
			i++ // Dummy value here
		}
		hasFilter := ints[i]
		i++
		var nFilters int
		if hasFilter == 1 {
			nFilters = ints[i]
			i++
		}
		c.Patterns[j].Filters = make([]Filter, nFilters)
		for k := 0; k < nFilters; k++ {
			c.Patterns[j].Filters[k].Volume = ints[i]
			i++
			c.Patterns[j].Filters[k].LPFCutoff = ints[i]
			i++
			c.Patterns[j].Filters[k].LPFResonance = ints[i]
			i++
		}
	}
	songLength := ints[i]
	i++
	c.LoopStart = ints[i]
	i++
	c.LoopEnd = ints[i]
	i++
	c.Arrangement = make([][8]int, songLength)
	for j := 0; j < songLength; j++ {
		for k := 0; k < 8; k++ {
			c.Arrangement[j][k] = ints[i]
			i++
		}
	}
	spew.Dump(c)
}
