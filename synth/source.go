package synth

import (
	"time"

	"github.com/200sc/klangsynthese/audio"
)

type Source struct {
	audio.Format
	Pitch   Pitch
	Volume  float64
	Seconds float64
}

func (s Source) Duration() time.Duration {
	return time.Duration(s.Seconds) * 1000 * time.Millisecond
}

func (s Source) Phase(i int) float64 {
	return phase(s.Pitch, i, s.SampleRate)
}

func (s Source) Update(opts ...Option) Source {
	for _, opt := range opts {
		s = opt(s)
	}
	return s
}

var (
	Int16 = Source{
		Format: audio.Format{
			44100,
			2,
			16,
		},
		Pitch:   A4,
		Volume:  .25,
		Seconds: 1,
	}
)
