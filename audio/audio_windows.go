//+build windows

package audio

import (
	"errors"
	"syscall"

	"github.com/oov/directsound-go/dsound"
)

var (
	user32           = syscall.NewLazyDLL("user32")
	getDesktopWindow = user32.NewProc("GetDesktopWindow")
	ds               *dsound.IDirectSound
	err              error
)

func init() {
	hasDefaultDevice := false
	dsound.DirectSoundEnumerate(func(guid *dsound.GUID, description string, module string) bool {
		if guid == nil {
			hasDefaultDevice = true
			return false
		}
		return true
	})
	if !hasDefaultDevice {
		ds = nil
		err = errors.New("No default device available to play audio off of")
		return
	}

	ds, err = dsound.DirectSoundCreate(nil)

	desktopWindow, _, err2 := getDesktopWindow.Call()
	if err != nil {
		ds = nil
		err = err2
		return
	}
	err = ds.SetCooperativeLevel(syscall.Handle(desktopWindow), dsound.DSSCL_PRIORITY)
	if err != nil {
		ds = nil
	}
}

func EncodeBytes(b []byte) (Audio, error) {
	if err != nil {
		return nil, err
	}

	// ???????????
	// ??????????
	// ?????????
	// this is for wav!
	// we need this to be universal!!!!
	// ?????????
	wf := dsound.WaveFormatEx{
		FormatTag:      dsound.WAVE_FORMAT_PCM,
		Channels:       Channels,
		SamplesPerSec:  SampleRate,
		BitsPerSample:  Bits,
		BlockAlign:     Channels * Bits / 8,
		AvgBytesPerSec: BytesPerSec,
		ExtSize:        0,
	}

	buffdsc := dsound.BufferDesc{
		// These flags cover everything we should ever want to do
		Flags:       dsound.DSBCAPS_GLOBALFOCUS | dsound.DSBCAPS_GETCURRENTPOSITION2 | dsound.DSBCAPS_CTRLVOLUME | dsound.DSBCAPS_CTRLPAN | dsound.DSBCAPS_CTRLFREQUENCY | dsound.DSBCAPS_LOCDEFER,
		Format:      &wf,
		BufferBytes: uint32(len(b)),
	}

	// Create the object which stores the wav data in a playable format
	dsbuff, err := ds.CreateSoundBuffer(&buffdsc)
	if err != nil {
		return nil, err
	}

	// Reserve some space in the sound buffer object to write to.
	// The Lock function (and by extension LockBytes) actually
	// reserves two spaces, but we ignore the second.
	by1, by2, err := dsbuff.LockBytes(0, uint32(len(b)), 0)
	if err != nil {
		return nil, err
	}

	// Write to the pointer we were given.
	copy(by1, b)

	// Update the buffer object with the new data.
	err = dsbuff.UnlockBytes(by1, by2)
	if err != nil {
		return nil, err
	}

	return &dsAudio{dsbuff, 0}, nil
}

type dsAudio struct {
	*dsound.IDirectSoundBuffer
	flags dsound.BufferPlayFlag
}

func (ds *dsAudio) Play() <-chan error {
	//... ???
	ch := make(chan error)
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

func (ds *dsAudio) Copy() *dsAudio {

}
