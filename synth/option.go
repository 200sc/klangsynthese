package synth

import "time"

type Option func(Source) Source

func Duration(t time.Duration) Option {
	return func(s Source) Source {
		s.Seconds = t.Seconds()
		return s
	}
}

func Volume(v float64) Option {
	return func(s Source) Source {
		if v > 1.0 {
			v = 1.0
		} else if v < 0 {
			v = 0
		}
		s.Volume = v
		return s
	}
}

func AtPitch(p Pitch) Option {
	return func(s Source) Source {
		s.Pitch = p
		return s
	}
}
