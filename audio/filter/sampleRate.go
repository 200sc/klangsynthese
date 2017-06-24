package filter

import "github.com/200sc/klangsynthese/audio"

type SampleRate func(*uint32)
type SupportsSampleRate interface {
	GetSampleRate() *uint32
}

func (srf SampleRate) Apply(a audio.Audio) audio.Audio {
	if ssr, ok := a.(SupportsSampleRate); ok {
		srf(ssr.GetSampleRate())
	}
	return a
}
