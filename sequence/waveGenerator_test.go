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
		Volumes(2000),
		Holds(time.Millisecond*150),
		HoldAt(time.Millisecond*400, 3),
		Ticks(time.Millisecond*200),
		Waves(synth.Int16.Sin, synth.Int16.Saw),
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
	loopsAndTicks := And(
		Ticks(time.Millisecond*200),
		Loops(true),
	)
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
			2000,
			2500,
			3000,
		),
		Waves(synth.Int16.Sin),
		loopsAndTicks,
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
		Volumes(500),
		Waves(
			synth.Int16.Square,
			synth.Int16.Square,
			synth.Int16.Sin,
			synth.Int16.Saw,
		),
		loopsAndTicks,
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
			2000,
			2500,
			3000,
		),
		Ticks(time.Millisecond*200),
		Waves(synth.Int16.Sin),
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
		Volumes(500),
		Ticks(time.Millisecond*200),
		Waves(
			synth.Int16.Square,
			synth.Int16.Square,
			synth.Int16.Sin,
			synth.Int16.Saw,
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
