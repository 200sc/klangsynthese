package sequence

import (
	"fmt"
	"testing"
	"time"

	"github.com/200sc/klangsynthese/synth"
)

func TestWaveGenerator(t *testing.T) {
	wg := WaveGenerator{
		// Todo: nicer syntax
		Fn: synth.Sin,
		// Todo: hold pattern, so individual notes can last longer
		// or shorter than the tick
		Tick: time.Millisecond * 200,
		PitchPattern: []synth.Pitch{
			synth.A4,
			synth.A5,
			synth.A6,
			synth.G6,
			synth.Rest,
			synth.G4,
		},
		VolumePattern: []synth.Volume{
			32,
			64,
			96,
		},
		Loop: true,
	}
	sq := wg.Generate()
	sq.Play()
	fmt.Println("Playing sequence")
	time.Sleep(5 * time.Second)
}

func TestCombineSeq(t *testing.T) {
	wg := WaveGenerator{
		// Todo: nicer syntax
		Fn: synth.Sin,
		// Todo: hold pattern, so individual notes can last longer
		// or shorter than the tick
		Tick: time.Millisecond * 200,
		PitchPattern: []synth.Pitch{
			synth.A4,
			synth.A5,
			synth.A6,
			synth.G6,
			synth.Rest,
			synth.G4,
		},
		VolumePattern: []synth.Volume{
			32,
			64,
			96,
		},
		Loop: true,
	}
	sq := wg.Generate()
	wg = WaveGenerator{
		// Todo: nicer syntax
		Fn: synth.Square,
		// Todo: hold pattern, so individual notes can last longer
		// or shorter than the tick
		Tick: time.Millisecond * 200,
		PitchPattern: []synth.Pitch{
			synth.C4,
			synth.C5,
			synth.C6,
			synth.C6,
			synth.C5,
			synth.C4,
		},
		VolumePattern: []synth.Volume{
			8,
		},
		Loop: true,
	}
	sq2 := wg.Generate()
	sq3, _ := sq.Combine(sq2)
	sq3.Play()
	fmt.Println("Playing sequence")
	time.Sleep(5 * time.Second)
}
