//+build darwin

package audio

import (
	"errors"

	"github.com/200sc/klangsynthese/audio/filter/supports"
)

type darwinNopAudio struct {
	Encoding
}

func (dna *darwinNopAudio) Play() <-chan error {
	ch := make(chan error)
	go func() {
		ch <- errors.New("Playback on Darwin is not supported")
	}()
	return ch
}

func (dna *darwinNopAudio) Stop() error {
	return errors.New("Playback on Darwin is not supported")
}

func (dna *darwinNopAudio) Filter(fs ...Filter) (Audio, error) {
	var a Audio = dna
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
	return dna, consError
}

func (dna *darwinNopAudio) MustFilter(fs ...Filter) Audio {
	a, _ := dna.Filter(fs...)
	return a
}

func EncodeBytes(enc Encoding) (Audio, error) {
	return &darwinNopAudio{enc}, nil
}
