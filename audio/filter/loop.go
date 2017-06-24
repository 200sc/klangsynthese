package filter

import "github.com/200sc/klangsynthese/audio"

type Loop func(*bool)
type SupportsLoop interface {
	GetLoop() *bool
}

func (lf Loop) Apply(a audio.Audio) audio.Audio {
	if sl, ok := a.(SupportsLoop); ok {
		lf(sl.GetLoop())
	}
	return a
}

func LoopOn(b *bool) {
	*b = true
}

func LoopOff(b *bool) {
	*b = false
}