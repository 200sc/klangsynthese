// Package flac provides functionality to handle .flac files and .flac encoded data
package flac

import (
	"errors"
	"fmt"
	"io"

	"github.com/200sc/klangsynthese/audio"
	"github.com/eaburns/flac"
)

// def wav format
var format = audio.Format{
	SampleRate: 44100,
	Bits:       16,
	Channels:   2,
}

// A Controller might eventually contain device information concerning
// what audios made from this should play out of but also might not
// exist in the future
type Controller struct{}

// NewController returns a default controller
func NewController() *Controller {
	return &Controller{}
}

// Wave encodes raw bytes with the default wavformatting into audio
// todo: this really shouldn't be here. Having some controller type that
// knows its format makes sense, but the output data has nothing to do with
// wav files.
func (mc *Controller) Wave(b []byte) (audio.Audio, error) {
	return audio.EncodeBytes(
		audio.Encoding{
			Data:   b,
			Format: mc.Format(),
		})
}

// Load loads wav data from the incoming reader as an audio
func (mc *Controller) Load(r io.Reader) (audio.Audio, error) {
	data, meta, err := flac.Decode(r)
	if err != nil {
		fmt.Println("Load error", err)
		return nil, err
	}

	fformat := audio.Format{
		SampleRate: uint32(meta.SampleRate),
		Channels:   uint16(meta.NChannels),
		Bits:       uint16(meta.BitsPerSample),
	}
	return audio.EncodeBytes(
		audio.Encoding{
			Data:   data,
			Format: fformat,
		})
}

// Save will eventually save an audio encoded as a wav to the given writer
func (mc *Controller) Save(r io.ReadWriter, a audio.Audio) error {
	return errors.New("Unsupported Functionality")
}

// Format returns the default wav formatting
func (mc *Controller) Format() audio.Format {
	return format
}
