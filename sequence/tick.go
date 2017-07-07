package sequence

import "time"

type Tick time.Duration

type HasTicks interface {
	GetTickPattern() time.Duration
	SetTickPattern(time.Duration)
}

func (vp *Tick) GetTickPattern() time.Duration {
	return time.Duration(*vp)
}

func (vp *Tick) SetTickPattern(vs time.Duration) {
	*vp = Tick(vs)
}

// Ticks sets the generator's Tick
func Ticks(t time.Duration) Option {
	return func(g Generator) {
		if ht, ok := g.(HasTicks); ok {
			ht.SetTickPattern(t)
		}
	}
}
