package sequence

import "time"

type HoldPattern []time.Duration

type HasHolds interface {
	GetHoldPattern() []time.Duration
	SetHoldPattern([]time.Duration)
}

func (hp *HoldPattern) GetHoldPattern() []time.Duration {
	return *hp
}

func (hp *HoldPattern) SetHoldPattern(ts []time.Duration) {
	*hp = ts
}

// Holds sets the generator's Hold pattern
func Holds(vs ...time.Duration) Option {
	return func(g Generator) {
		if hhs, ok := g.(HasHolds); ok {
			hhs.SetHoldPattern(vs)
		}
	}
}

// HoldAt sets the n'th value in the entire play sequence
// to be Hold p. This could involve duplicating a pattern
// until it is long enough to reach n. Meaningless if the
// Hold pattern has not been set yet.
func HoldAt(t time.Duration, n int) Option {
	return func(g Generator) {
		if hhs, ok := g.(HasHolds); ok {
			if hl, ok := hhs.(HasLength); ok {
				if hl.GetLength() < n {
					Holds := hhs.GetHoldPattern()
					if len(Holds) == 0 {
						return
					}
					// If the pattern is not long enough, there are two things
					// we could do-- 1. Extend the pattern and replace the
					// individual note, or 2. Replace the note that would be
					// played at n and thus all earlier and later plays within
					// the pattern as well.
					//
					// This uses approach 1.
					for len(Holds) <= n {
						Holds = append(Holds, Holds...)
					}
					Holds[n] = t
					hhs.SetHoldPattern(Holds)
				}
			}
		}
	}
}

// HoldPatternAt sets the n'th value in the Hold pattern
// to be Hold p. Meaningless if the Hold pattern has not
// been set yet.
func HoldPatternAt(t time.Duration, n int) Option {
	return func(g Generator) {
		if hhs, ok := g.(HasHolds); ok {
			Holds := hhs.GetHoldPattern()
			if len(Holds) <= n {
				return
			}
			Holds[n] = t
			hhs.SetHoldPattern(Holds)
		}
	}
}
