//+build windows

package audio

import (
	"github.com/200sc/klangsynthese/audio/filter/supports"
	"github.com/oov/directsound-go/dsound"
)

type dsAudio struct {
	*Encoding
	*dsound.IDirectSoundBuffer
	flags dsound.BufferPlayFlag
	pan   float64
}

func (ds *dsAudio) Play() <-chan error {
	ch := make(chan error)
	if ds.Loop {
		ds.flags = dsound.DSBPLAY_LOOPING
	}
	go func(dsbuff *dsound.IDirectSoundBuffer, flags dsound.BufferPlayFlag, ch chan error) {
		err := dsbuff.SetCurrentPosition(0)
		if err != nil {
			ch <- err
		} else {
			err = dsbuff.Play(0, flags)
			if err != nil {
				ch <- err
			} else {
				ch <- nil
			}
		}
	}(ds.IDirectSoundBuffer, ds.flags, ch)
	return ch
}

func (ds *dsAudio) Stop() error {
	err := ds.IDirectSoundBuffer.Stop()
	if err != nil {
		return err
	}
	return ds.IDirectSoundBuffer.SetCurrentPosition(0)
}

func (ds *dsAudio) Filter(fs ...Filter) (Audio, error) {
	var a Audio = ds
	var err error
	var consError supports.ConsError
	for _, f := range fs {
		a, err = f.Apply(a)
		if err != nil {
			if consError == nil {
				consError = err.(supports.ConsError)
			} else {
				consError = consError.Cons(err)
			}
		}
	}
	pan := ds.pan
	// Consider: this is a significant amount
	// of work to do just to make this an in-place filter.
	// would it be worth it to offer both in place and non-inplace
	// filter functions?
	a2, err := EncodeBytes(*ds.Encoding)
	if err != nil {
		return nil, err
	}
	// reassign the contents of ds to be that of the
	// new audio, so that this filters in place
	*ds = *a2.(*dsAudio)
	ds.SetPan(pan)
	return ds, consError
}

// MustFilter acts like Filter, but ignores errors (it does not panic,
// as filter errors are expected to be non-fatal)
func (ds *dsAudio) MustFilter(fs ...Filter) Audio {
	var a Audio = ds
	for _, f := range fs {
		a, _ = f.Apply(a)
	}
	return a
}

var (
	RIGHT_PAN = 10000.0
	LEFT_PAN  = -10000.0
)

func (ds *dsAudio) GetPan() float64 {
	pan, err := ds.IDirectSoundBuffer.GetPan()
	if err != nil {
		// This ID might need to return float64, error
		return 0
	}
	// Scale the pan value by what we know about dsound
	return float64(pan) / RIGHT_PAN
}

func (ds *dsAudio) SetPan(pan float64) {
	ds.pan = pan
	// As in GetPan, returns an error that we should probably return
	ds.IDirectSoundBuffer.SetPan(int32(pan * RIGHT_PAN))
}

func (ds *dsAudio) Copy() (Audio, error) {
	a, err := ds.Encoding.Copy()
	if err != nil {
		return a, err
	}
	ds2 := a.(*dsAudio)
	ds2.SetPan(ds.pan)
	return ds2, nil
}

func (ds *dsAudio) MustCopy() Audio {
	a, err := ds.Copy()
	if err != nil {
		panic(err)
	}
	return a
}
