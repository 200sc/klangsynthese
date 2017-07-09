package filter

import (
	"github.com/200sc/klangsynthese/audio"
	"github.com/200sc/klangsynthese/audio/filter/supports"
)

// A SampleRate is a function that takes in uint32 SampleRates
type SampleRate func(*uint32)

// Apply checks that the given audio supports SampleRate, filters if it
// can, then returns
func (srf SampleRate) Apply(a audio.Audio) (audio.Audio, error) {
	if ssr, ok := a.(supports.SampleRate); ok {
		srf(ssr.GetSampleRate())
		return a, nil
	}
	return a, supports.NewUnsupported([]string{"SampleRate"})
}

// Speed might slow down or speed up a sample, but this will probably
// just effect the perceived pitch of the sample, and a more complicated
// method is needed to modify speed without modifying perceived pitch
func Speed(mult float64) SampleRate {
	return func(sr *uint32) {
		*sr = uint32(float64(*sr) * mult)
	}
}
