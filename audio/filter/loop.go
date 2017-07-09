package filter

import (
	"github.com/200sc/klangsynthese/audio"
	"github.com/200sc/klangsynthese/audio/filter/supports"
)

// Loop functions modify a boolean, with the intention that that boolean
// is a loop variable
type Loop func(*bool)

// Apply checks that the given audio supports Loop, filters if it
// can, then returns
func (lf Loop) Apply(a audio.Audio) (audio.Audio, error) {
	if sl, ok := a.(supports.Loop); ok {
		lf(sl.GetLoop())
		return a, nil
	}
	return a, supports.NewUnsupported([]string{"Loop"})
}

// LoopOn sets the loop to hapen
func LoopOn() Loop {
	return func(b *bool) {
		*b = true
	}
}

// LoopOff sets the loop to not happen
func LoopOff() Loop {
	return func(b *bool) {
		*b = false
	}
}
