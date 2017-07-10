package sequence

import (
	"errors"
	"time"

	"github.com/200sc/klangsynthese/audio"
)

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
	Pattern      []*audio.Multi
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
				s.Pattern[s.patternIndex].Play()
				select {
				case <-s.stopCh:
					s.stopCh <- s.Pattern[s.patternIndex].Stop()
					ch <- nil
					return
				case <-s.Ticker.C:
				}
				s.patternIndex++
			}
			if !s.loop {
				ch <- nil
				return
			}
		}
	}()
	return ch
}

func (s *Sequence) Filter(fs ...audio.Filter) (audio.Audio, error) {
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
	return s, nil
}

func (s *Sequence) MustFilter(fs ...audio.Filter) audio.Audio {
	a, _ := s.Filter(fs...)
	return a
}

func (s *Sequence) Stop() error {
	s.stopCh <- nil
	return <-s.stopCh
}

func (s *Sequence) Copy() (audio.Audio, error) {
	var err error
	s2 := &Sequence{
		Pattern:      make([]*audio.Multi, len(s.Pattern)),
		Ticker:       time.NewTicker(s.tickDuration),
		tickDuration: s.tickDuration,
		stopCh:       make(chan error),
		loop:         s.loop,
	}
	for i := range s2.Pattern {
		s2.Pattern[i] = new(audio.Multi)
		s2.Pattern[i].Audios = make([]audio.Audio, len(s.Pattern[i].Audios))
		for j := range s2.Pattern[i].Audios {
			// This could make a sequence that reuses the same
			// audio use a lot more memory when copied-- a better route
			// would involve identifying all unique audios
			// and making a copy for each of those, but that
			// requires producing unique IDs for each audio
			// (which would probably be a hash of their encoding?
			// but that raises issues for audios that don't want
			// to follow real encoding rules (like this one!))
			s2.Pattern[i].Audios[j], err = s.Pattern[i].Audios[j].Copy()
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

func (s *Sequence) Mix(s2 *Sequence) (*Sequence, error) {
	// Todo: we should be able to combine not-too-disparate
	// sequences like one that ticks on .5 seconds and one that ticks
	// on .25 seconds
	if s.tickDuration != s2.tickDuration {
		return nil, errors.New("Incompatible sequences")
	}
	seq, err := s.Copy()
	if err != nil {
		return nil, err
	}
	s3 := seq.(*Sequence)
	for i, col := range s2.Pattern {
		s3.Pattern[i].Audios = append(s3.Pattern[i].Audios, col.Audios...)
	}
	return s3, nil
}

func (s *Sequence) Append(s2 *Sequence) (*Sequence, error) {
	// Todo: we should be able to combine not-too-disparate
	// sequences like one that ticks on .5 seconds and one that ticks
	// on .25 seconds
	if s.tickDuration != s2.tickDuration {
		return nil, errors.New("Incompatible sequences")
	}
	seq, err := s.Copy()
	if err != nil {
		return nil, err
	}
	s3 := seq.(*Sequence)
	s3.Pattern = append(s3.Pattern, s2.Pattern...)
	return s3, nil
}
