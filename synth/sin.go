package synth

import "math"

const (
	SampleRate = 44100
)

// Todo: functions to make frequency and volume meaningful!
func Sin(freq uint16, seconds float64, volume uint8) []byte {
	freqf := float64(freq / 4)
	wave := make([]byte, int(seconds*float64(SampleRate)))
	for i := range wave {
		wave[i] = byte(float64(volume) * math.Sin(freqf*(float64(i)/float64(SampleRate))*2*math.Pi))
	}
	return wave
}
