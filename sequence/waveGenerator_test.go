package sequence

import (
	"fmt"
	"testing"
	"time"

	"github.com/200sc/klangsynthese/synth"
)

func TestWaveGenerator(t *testing.T) {
	wg := WaveGenerator{
		Fn:   synth.Sin,
		Tick: time.Millisecond * 500,
		PitchPattern: []synth.Pitch{
			synth.A4,
			synth.A5,
			synth.A6,
			synth.G6,
			synth.G5,
			synth.G4,
		},
		VolumePattern: []synth.Volume{
			32,
			64,
			96,
		},
		Loop: false,
	}
	sq := wg.Generate()
	sq.Play()
	fmt.Println("Playing sequence")
	time.Sleep(5 * time.Second)
}
