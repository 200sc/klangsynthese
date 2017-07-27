package audio

import (
	"time"

	"github.com/pkg/errors"
)

// A Multi lets lists of audios be used simultaneously
type Multi struct {
	Audios []Audio
}

// NewMulti returns a new multi
func NewMulti(as ...Audio) *Multi {
	return &Multi{Audios: as}
}

// Play plays all audios in the Multi ASAP
func (m *Multi) Play() <-chan error {
	extCh := make(chan error)
	go func() {
		// Todo: Propagating N errors?
		for _, a := range m.Audios {
			a.Play()
		}
		extCh <- nil
	}()
	return extCh
}

// Filter applies all the given filters on everything in the Multi
func (m *Multi) Filter(fs ...Filter) (Audio, error) {
	var err error
	for i, a := range m.Audios {
		m.Audios[i], err = a.Filter(fs...)
		err = errors.Wrap(err, "Failed to apply filter")
	}
	return m, err
}

// MustFilter acts like filter but ignores errors.
func (m *Multi) MustFilter(fs ...Filter) Audio {
	a, _ := m.Filter(fs...)
	return a
}

// Stop stops all audios in the Multi. Any that fail will report an error.
func (m *Multi) Stop() error {
	var err error
	for _, a := range m.Audios {
		err = errors.Wrap(a.Stop(), "Failed to stop audio")
	}
	return err
}

// Copy returns a copy of this Multi
func (m *Multi) Copy() (Audio, error) {
	var err error
	newAudios := make([]Audio, len(m.Audios))
	for i, a := range m.Audios {
		newAudios[i], err = a.Copy()
		if err != nil {
			return nil, err
		}
	}
	return &Multi{newAudios}, nil

}

// MustCopy acts like Copy but panics if error != nil
func (m *Multi) MustCopy() Audio {
	m2, err := m.Copy()
	if err != nil {
		panic(err)
	}
	return m2
}

// PlayLength returns how long this audio will play for
func (m *Multi) PlayLength() time.Duration {
	var d time.Duration
	for _, a := range m.Audios {
		d2 := a.PlayLength()
		if d < d2 {
			d = d2
		}
	}
	return d
}
