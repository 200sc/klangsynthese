package filter

import (
	"fmt"

	"github.com/200sc/klangsynthese/audio/filter/supports"
)

func Resample(newSampleRate uint32, pitchShifter PitchShifter) Encoding {
	return func(senc supports.Encoding) {
		oldSampleRate := *senc.GetSampleRate()
		ratio := float64(newSampleRate) / float64(oldSampleRate)
		Speed(ratio, pitchShifter)(senc)
		fmt.Println(newSampleRate, *senc.GetSampleRate())
	}
}

func Speed(ratio float64, pitchShifter PitchShifter) Encoding {
	return func(senc supports.Encoding) {
		r := ratio
		fmt.Println(ratio)
		for r < .5 {
			r *= 2
			pitchShifter.PitchShift(.5)(senc)
		}
		for r > 2.0 {
			r /= 2
			pitchShifter.PitchShift(2.0)(senc)
		}
		pitchShifter.PitchShift(1 / r)(senc)
		ModSampleRate(ratio)(senc.GetSampleRate())
	}
}
