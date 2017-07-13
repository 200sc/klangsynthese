package ceol

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/200sc/klangsynthese/sequence"
	"github.com/200sc/klangsynthese/synth"
	"github.com/stretchr/testify/assert"
)

func TestReadCeol(t *testing.T) {
	f, err := os.Open("test.ceol")
	assert.Nil(t, err)
	c, err := Open(f)
	assert.Nil(t, err)
	wg := sequence.NewWaveGenerator(
		sequence.Chords(c.ChordPattern()),
		sequence.Volumes(2000),
		sequence.Ticks(DurationFromQuarters(c.Bpm, 1)),
		sequence.Waves(synth.Int16.Saw),
	)
	sq := wg.Generate()
	sq.Play()
	fmt.Println("Playing sequence")
	time.Sleep(5 * time.Second)
	sq.Stop()
}
