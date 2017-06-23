package sequence

import "github.com/200sc/klangsynthese/synth"

type WaveGenerator struct {
	Fn           synth.Wave
	PitchPattern []synth.Pitch
	// seconds between waves
	SecondPattern []float64
	VolumePattern []synth.Volume
}

func (wg *WaveGenerator) Generate() *Sequence {
	
}
