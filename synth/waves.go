package synth

import (
	"fmt"
	"math"
)

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

func Square(freq uint16, seconds float64, volume uint8) []byte {
	wave := Sin(freq, seconds, volume)
	for i, v := range wave {
		if v < 128 {
			wave[i] = volume - 1
		} else {
			wave[i] = 255 - (volume) + 2
		}
	}
	fmt.Println(wave)
	return wave
}
