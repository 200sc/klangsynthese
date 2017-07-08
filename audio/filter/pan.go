package filter

import (
	"github.com/200sc/klangsynthese/audio"
	"github.com/200sc/klangsynthese/audio/filter/supports"
)

// Pan filters are functions on float64s
// the expected input for a pan function should be
// -1 -> left, 0 -> center, +1 -> right,
type Pan func(*float64)

// Apply checks that the given audio supports Pan, filters if it
// can, then returns
func (p Pan) Apply(a audio.Audio) (audio.Audio, error) {
	if sp, ok := a.(supports.Pan); ok {
		f := sp.GetPan()
		p(&f)
		sp.SetPan(f)
		return a, nil
	}
	return a, supports.NewUnsupported([]string{"Pan"})
}

func RightPan() Pan {
	return func(f *float64) {
		*f = 1
	}
}

func LeftPan() Pan {
	return func(f *float64) {
		*f = -1
	}
}

func CenterPan() Pan {
	return func(f *float64) {
		*f = 0
	}
}
