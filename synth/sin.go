package synth

import "math"

func Sin(freq, seconds float64, volume uint8, sampleRate uint32) []byte {
	wave := make([]byte, int(seconds*float64(sampleRate)))
	for i := range wave {
		wave[i] = byte(float64(volume) * math.Sin(freq*(float64(i)/float64(sampleRate))*2*math.Pi))
	}
	return wave
}
