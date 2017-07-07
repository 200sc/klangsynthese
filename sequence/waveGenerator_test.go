package sequence

import (
	"fmt"
	"testing"
	"time"

	"github.com/200sc/klangsynthese/synth"
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
	sq.Play()
	fmt.Println("Playing sequence")
	time.Sleep(5 * time.Second)
}

func TestCombineSeq(t *testing.T) {
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
		Waves(synth.Square),
		Loops(true),
	)
	sq2 := wg.Generate()
	sq3, _ := sq.Combine(sq2)
	sq3.Play()
	fmt.Println("Playing sequence")
	time.Sleep(5 * time.Second)
}
