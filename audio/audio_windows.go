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
	return a.MustCopy(), consError
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
