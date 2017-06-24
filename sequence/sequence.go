package sequence

import (
	"time"

	"github.com/200sc/klangsynthese/audio"
)

type SequenceI interface {
	audio.Audio
}

// This is notes / pseudo-code / not useable yet

// A Sequence does not care if it loops because that is audio/Encoding's job
// A Sequence does not care how long it should play each sample it is given
// because that is the job of the individual samples
// A Sequence does care how much time it should wait between samples
// A Sequence does care if that time is variable (swing rhythm)
// A Sequence satisfies Audio
type Sequence struct {
	// Sequences play patterns of audio
	// everything at Pattern[0] will be simultaneously Play()ed at
	// Sequence.Play()
	Pattern      [][]audio.Audio
	patternIndex int
	// Every tick, the next index in Pattern will be played by a Sequence
	// until the pattern is over.
	Ticker *time.Ticker
	// needed to copy Ticker
	// consider: replacing ticker with dynamic ticker
	tickDuration time.Duration
	stopCh       chan error
	loop         bool
}

func (s *Sequence) Play() <-chan error {
	ch := make(chan error)
	go func() {
		for {
			s.patternIndex = 0
			for s.patternIndex < len(s.Pattern) {
				for _, a := range s.Pattern[s.patternIndex] {
					a.Play()
				}
				select {
				case <-s.stopCh:
					for _, a := range s.Pattern[s.patternIndex] {
						err := a.Stop()
						if err != nil {
							s.stopCh <- err
							return
						}
					}
					return
				case <-s.Ticker.C:
				}
				s.patternIndex++
			}
			if !s.loop {
				return
			}
		}
	}()
	return ch
}

func (s *Sequence) Filter(fs ...audio.Filter) audio.Audio {
	// Filter on a sequence just applies the filter to all audios..
	// but it can't do that always, what if the filter is Loop?
	// this implies two kinds of filters?
	// this doesn't work because FIlter is not an interface
	// for _, f := range fs {
	// 	if _, ok := f.(audio.Loop); ok {
	// 		s.loop = true
	// 	} else if _, ok := f.(audio.NoLoop); ok {
	// 		s.loop = false
	// 	} else {
	// 		for _, col := range s.Pattern {
	// 			for _, a := range col {
	// 				a.Filter(f)
	// 			}
	// 		}
	// 	}
	// }
	return s
}

func (s *Sequence) Stop() error {
	s.stopCh <- nil
	return <-s.stopCh
}

func (s *Sequence) Copy() (audio.Audio, error) {
	var err error
	s2 := &Sequence{
		Pattern:      make([][]audio.Audio, len(s.Pattern)),
		Ticker:       time.NewTicker(s.tickDuration),
		tickDuration: s.tickDuration,
		stopCh:       make(chan error),
		loop:         s.loop,
	}
	for i := range s2.Pattern {
		s2.Pattern[i] = make([]audio.Audio, len(s.Pattern[i]))
		for j := range s2.Pattern[i] {
			// This could make a sequence that reuses the same
			// audio use a lot more memory when copied-- a better route
			// would involve identifying all unique audios
			// and making a copy for each of those, but that
			// requires producing unique IDs for each audio
			// (which would probably be a hash of their encoding?
			// but that raises issues for audios that don't want
			// to follow real encoding rules (like this one!))
			s2.Pattern[i][j], err = s.Pattern[i][j].Copy()
			if err != nil {
				return nil, err
			}
		}
	}
	return s2, nil
}

func (s *Sequence) MustCopy() audio.Audio {
	a, err := s.Copy()
	if err != nil {
		panic(err)
	}
	return a
}

func (s *Sequence) GetEncoding() *audio.Encoding {
	// With the interface the way it is now, this would require
	// combining all encodings in a sequence and if there were
	// conflicts in formatting we would have to just pick one
	// this is a lot of work, and causes issues, so we aren't
	// comitting to that right yet
	return nil
}
