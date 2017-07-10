package audio

import "github.com/200sc/klangsynthese/audio/filter/supports"

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
	var consError supports.ConsError
	for i, a := range m.Audios {
		m.Audios[i], err = a.Filter(fs...)
		if err != nil {
			if consError == nil {
				consError = err.(supports.ConsError)
			} else {
				consError = consError.Cons(err)
			}
		}
	}
	return m, consError
}

// MustFilter acts like filter but ignores errors.
func (m *Multi) MustFilter(fs ...Filter) Audio {
	a, _ := m.Filter(fs...)
	return a
}

// Stop stops all audios in the Multi. If any of the audios fail to be stopped,
// following audios will not try to stop. Consider: should this behavior change?
func (m *Multi) Stop() error {
	for _, a := range m.Audios {
		err := a.Stop()
		// Todo: consError?
		if err != nil {
			return err
		}
	}
	return nil
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
