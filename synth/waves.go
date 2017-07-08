package synth

import "math"

// Default wave generation variables
const (
	SampleRate = 44100
	Channels   = 2
)

// Thanks to https://en.wikibooks.org/wiki/Sound_Synthesis_Theory/Oscillators_and_Wavetables
func phase(freq Pitch, i, sampleRate int) float64 {
	return float64(freq/4) * (float64(i) / float64(SampleRate)) * 2 * math.Pi
}

// A Wave function takes in a volume, pitch, and time and produces a sound wave
type Wave func(Pitch, float64, Volume) []byte

// Sin produces a Sin wave
//         __
//       --  --
//      /      \
//--__--        --__--
func Sin(freq Pitch, seconds float64, volume Volume) []byte {
	wave := make([]byte, int(seconds*float64(SampleRate)*Channels))
	for i := 0; i < len(wave); i += 4 {
		// Todo: helper to do this for every wave
		val := int(float64(volume) * math.Sin(phase(freq, i, SampleRate)))
		wave[i] = byte(val % 256)
		wave[i+1] = byte(val >> 8)
		wave[i+2] = wave[i]
		wave[i+3] = wave[i+1]
	}
	return wave
}

// Pulse acts like Square when given a pulse of 2, when given any lesser
// pulse the time up and down will change so that 1/pulse time the wave will
// be up.
//
//     __    __
//     ||    ||
// ____||____||____
func Pulse(pulse float64) Wave {
	pulseSwitch := 1 - 2/pulse
	return func(freq Pitch, seconds float64, volume Volume) []byte {
		wave := make([]byte, int(seconds*float64(SampleRate)*Channels))
		for i := range wave {
			// alternatively phase % 2pi
			if math.Sin(phase(freq, i, SampleRate)) > pulseSwitch {
				wave[i] = byte(volume)
			} else {
				wave[i] = byte(-volume)
			}
		}
		return wave
	}
}

// Square produces a Square wave
//
//       _________
//       |       |
// ______|       |________
func Square(freq Pitch, seconds float64, volume Volume) []byte {
	wave := make([]byte, int(seconds*float64(SampleRate)*Channels))
	for i := range wave {
		// alternatively phase % 2pi
		if math.Sin(phase(freq, i, SampleRate)) > 0 {
			wave[i] = byte(volume)
		} else {
			wave[i] = byte(-volume)
		}
	}
	return wave
}

// Saw produces a saw wave
//
//   ^   ^   ^
//  / | / | /
// /  |/  |/
func Saw(freq Pitch, seconds float64, volume Volume) []byte {
	wave := make([]byte, int(seconds*float64(SampleRate)*Channels))
	for i := range wave {
		wave[i] = byte(float64(volume) - (float64(volume) / math.Pi * math.Mod(phase(freq, i, SampleRate), 2*math.Pi)))
	}
	return wave
}

// Triangle produces a Triangle wave
//
//   ^   ^
//  / \ / \
// v   v   v
func Triangle(freq Pitch, seconds float64, volume Volume) []byte {
	wave := make([]byte, int(seconds*float64(SampleRate)*2))
	for i := range wave {
		p := math.Mod(phase(freq, i, SampleRate), 2*math.Pi)
		m := byte(p * (2 * float64(volume) / math.Pi))
		if math.Sin(p) > 0 {
			wave[i] = byte(-volume) + m
		} else {
			wave[i] = 3*byte(volume) - m
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
