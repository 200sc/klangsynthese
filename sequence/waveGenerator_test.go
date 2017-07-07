package sequence

import (
	"fmt"
	"testing"
	"time"

	"github.com/200sc/klangsynthese/synth"
	"github.com/stretchr/testify/assert"
)

func TestWaveGenerator(t *testing.T) {
	wg := NewWaveGenerator(
		Pitches(
			synth.A4,
			synth.A5,
			synth.A6,
			synth.G6,
			synth.Rest,
			synth.G4,
		),
		Volumes(16),
		Holds(time.Millisecond*150),
		HoldAt(time.Millisecond*400, 3),
		Ticks(time.Millisecond*200),
		Waves(synth.Sin, synth.Saw),
		Loops(true),
		PlayLength(7),
	)
	sq := wg.Generate()
	sq.Play()
	fmt.Println("Playing sequence")
	time.Sleep(5 * time.Second)
	sq.Stop()
}

func TestMixSeq(t *testing.T) {
	wg := NewWaveGenerator(
		Pitches(
			synth.A4,
			synth.A5,
			synth.A6,
			synth.G6,
			synth.Rest,
			synth.G4,
		),
		Volumes(
			16,
			32,
			48,
		),
		Ticks(time.Millisecond*200),
		Waves(synth.Sin),
		Loops(true),
	)
	sq := wg.Generate()
	wg = NewWaveGenerator(
		Pitches(
			synth.C4,
			synth.C5,
			synth.C6,
			synth.C6,
			synth.C5,
			synth.C4,
		),
		Volumes(8),
		Ticks(time.Millisecond*200),
		Waves(
			synth.Square,
			synth.Square,
			synth.Sin,
			synth.Saw,
		),
		Loops(true),
	)
	sq2 := wg.Generate()
	sq3, err := sq.Mix(sq2)
	assert.Nil(t, err)
	sq3.Play()
	fmt.Println("Playing sequence")
	time.Sleep(5 * time.Second)
	assert.Nil(t, sq3.Stop())
}

func TestAppendSeq(t *testing.T) {
	wg := NewWaveGenerator(
		Pitches(
			synth.A4,
			synth.A5,
			synth.A6,
			synth.G6,
			synth.Rest,
			synth.G4,
		),
		Volumes(
			16,
			32,
			48,
		),
		Ticks(time.Millisecond*200),
		Waves(synth.Sin),
		Loops(true),
	)
	sq := wg.Generate()
	wg = NewWaveGenerator(
		Pitches(
			synth.C4,
			synth.C5,
			synth.C6,
			synth.C6,
			synth.C5,
			synth.C4,
		),
		Volumes(8),
		Ticks(time.Millisecond*200),
		Waves(
			synth.Square,
			synth.Square,
			synth.Sin,
			synth.Saw,
		),
		Loops(true),
	)
	sq2 := wg.Generate()
	sq3, err := sq2.Append(sq)
	assert.Nil(t, err)
	sq3.Play()
	fmt.Println("Playing sequence")
	time.Sleep(5 * time.Second)
	assert.Nil(t, sq3.Stop())
}
