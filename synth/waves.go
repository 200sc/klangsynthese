package synth

import "math"

const (
	SampleRate = 44100
)

// Thanks to https://en.wikibooks.org/wiki/Sound_Synthesis_Theory/Oscillators_and_Wavetables
func phase(freq Pitch, i, sampleRate int) float64 {
	return float64(freq/4) * (float64(i) / float64(SampleRate)) * 2 * math.Pi
}

type Wave func(Pitch, float64, uint8) []byte

func Sin(freq Pitch, seconds float64, volume uint8) []byte {
	wave := make([]byte, int(seconds*float64(SampleRate)))
	for i := range wave {
		wave[i] = byte(float64(volume) * math.Sin(phase(freq, i, SampleRate)))
	}
	return wave
}

// Pulse acts like Square when given a pulse of 2, when given any lesser
// pulse the time up and down will change so that 1/pulse time the wave will
// be up.
func Pulse(pulse float64) Wave {
	pulseSwitch := 1 - 2/pulse
	return func(freq Pitch, seconds float64, volume uint8) []byte {
		wave := make([]byte, int(seconds*float64(SampleRate)))
		for i := range wave {
			// alternatively phase % 2pi
			if math.Sin(phase(freq, i, SampleRate)) > pulseSwitch {
				wave[i] = volume
			} else {
				wave[i] = -volume
			}
		}
		return wave
	}
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

// Could have pulse triangle

// Reverse is included here so Reverse(Saw(...)) and the like can be written
func Reverse(wave []byte) []byte {
	for i := 0; i < len(wave)/2; i++ {
		j := len(wave) - i - 1
		wave[i], wave[j] = wave[j], wave[i]
	}
	return wave
}

// Add is a utility to add together all indices in the given waves.
// You need to send waves of the same length to Add, right now Add does
// not check that you are doing that. (strictly speaking, the first wave
// just needs to be as long as the longest following wave, all remaining
// waves can be of a shorter length than the first)
func Add(waves ...[]byte) []byte {
	wave := waves[0]
	for j := 1; j < len(waves); j++ {
		for i, v := range waves[j] {
			wave[i] += v
		}
	}
	return wave
}
