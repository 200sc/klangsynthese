package synth

import "math"

const (
	SampleRate = 44100
)

// Thanks to https://en.wikibooks.org/wiki/Sound_Synthesis_Theory/Oscillators_and_Wavetables
func phase(freq Pitch, i, sampleRate int) float64 {
	return float64(freq/4) * (float64(i) / float64(SampleRate)) * 2 * math.Pi
}

func Sin(freq Pitch, seconds float64, volume uint8) []byte {
	wave := make([]byte, int(seconds*float64(SampleRate)))
	for i := range wave {
		wave[i] = byte(float64(volume) * math.Sin(phase(freq, i, SampleRate)))
	}
	return wave
}

func Square(freq Pitch, seconds float64, volume uint8) []byte {
	wave := make([]byte, int(seconds*float64(SampleRate)))
	for i := range wave {
		// alternatively phase % 2pi
		if math.Sin(phase(freq, i, SampleRate)) > 0 {
			wave[i] = volume
		} else {
			wave[i] = -volume
		}
	}
	return wave
}

func Saw(freq Pitch, seconds float64, volume uint8) []byte {
	wave := make([]byte, int(seconds*float64(SampleRate)))
	for i := range wave {
		wave[i] = byte(float64(volume) - (float64(volume) / math.Pi * math.Mod(phase(freq, i, SampleRate), 2*math.Pi)))
	}
	return wave
}

func Triangle(freq Pitch, seconds float64, volume uint8) []byte {
	wave := make([]byte, int(seconds*float64(SampleRate)))
	for i := range wave {
		p := math.Mod(phase(freq, i, SampleRate), 2*math.Pi)
		m := byte(p * (2 * float64(volume) / math.Pi))
		if math.Sin(p) > 0 {
			wave[i] = -volume + m
		} else {
			wave[i] = 3*volume - m
		}
	}
	return wave
}

// Reverse is included here so Reverse(Saw(...)) and the like can be written
func Reverse(wave []byte) []byte {
	for i := 0; i < len(wave)/2; i++ {
		j := len(wave) - i - 1
		wave[i], wave[j] = wave[j], wave[i]
	}
	return wave
}

type Wave func(freq Pitch, i, sampleRate int)
