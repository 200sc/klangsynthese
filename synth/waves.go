package synth

import "math"

// Default wave generation variables
// Todo: add controllers for wave generation that have a built in format
// and bit depth, right now all these waves are channel = 2, bitdepth = 16, sampleRate = 44100
const (
	SampleRate = 44100
	Channels   = 2
)

// Thanks to https://en.wikibooks.org/wiki/Sound_Synthesis_Theory/Oscillators_and_Wavetables
func phase(freq Pitch, i, sampleRate int) float64 {
	return float64(freq) * (float64(i) / float64(SampleRate)) * 2 * math.Pi
}

func bytesFromInts(is []int16, channels int) []byte {
	wave := make([]byte, len(is)*channels*2)
	for i := 0; i < len(wave); i += channels * 2 {
		wave[i] = byte(is[i/4] % 256)
		wave[i+1] = byte(is[i/4] >> 8)
		// duplicate the contents across all channels
		for c := 1; c < channels; c++ {
			wave[i+(2*c)] = wave[i]
			wave[i+(2*c)+1] = wave[i+1]
		}
	}
	//fmt.Println(is)
	//fmt.Println(wave)
	return wave
}

// A Wave function takes in a volume, pitch, and time and produces a sound wave
type Wave func(Pitch, float64, Volume) []byte

// Sin produces a Sin wave
//         __
//       --  --
//      /      \
//--__--        --__--
func Sin(freq Pitch, seconds float64, volume Volume) []byte {
	wave := make([]int16, int(seconds*float64(SampleRate)))
	for i := 0; i < len(wave); i++ {
		wave[i] = int16(float64(volume) * math.Sin(phase(freq, i, SampleRate)))
	}
	return bytesFromInts(wave, Channels)
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
		wave := make([]int16, int(seconds*float64(SampleRate)))
		for i := range wave {
			// alternatively phase % 2pi
			if math.Sin(phase(freq, i, SampleRate)) > pulseSwitch {
				wave[i] = int16(volume)
			} else {
				wave[i] = int16(-volume)
			}
		}
		return bytesFromInts(wave, Channels)
	}
}

// Square produces a Square wave
//
//       _________
//       |       |
// ______|       |________
var Square = Pulse(2)

// Saw produces a saw wave
//
//   ^   ^   ^
//  / | / | /
// /  |/  |/
func Saw(freq Pitch, seconds float64, volume Volume) []byte {
	wave := make([]int16, int(seconds*float64(SampleRate)*Channels))
	for i := range wave {
		wave[i] = int16(float64(volume) - (float64(volume) / math.Pi * math.Mod(phase(freq, i, SampleRate), 2*math.Pi)))
	}
	return bytesFromInts(wave, Channels)
}

// Triangle produces a Triangle wave
//
//   ^   ^
//  / \ / \
// v   v   v
func Triangle(freq Pitch, seconds float64, volume Volume) []byte {
	wave := make([]int16, int(seconds*float64(SampleRate)*2))
	for i := range wave {
		p := math.Mod(phase(freq, i, SampleRate), 2*math.Pi)
		m := int16(p * (2 * float64(volume) / math.Pi))
		if math.Sin(p) > 0 {
			wave[i] = int16(-volume) + m
		} else {
			wave[i] = 3*int16(volume) - m
		}
	}
	return bytesFromInts(wave, Channels)
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
