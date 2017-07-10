// Package mp3 provides functionality to handle .mp3 files and .mp3 encoded data
package mp3

import (
	"bytes"
	"errors"
	"io"

	"github.com/200sc/klangsynthese/audio"

	haj "github.com/hajimehoshi/go-mp3"
)

var format = audio.Format{
	SampleRate: 44100,
	Bits:       16,
	Channels:   2,
}

// A Controller might eventually contain device information concerning
// what audios made from this should play out of but also might not
// exist in the future
type Controller struct{}

// NewController returns a new mp3 controller
func NewController() *Controller {
	return &Controller{}
}

// Load loads an mp3-encoded reader into an audio
func (mc *Controller) Load(r io.ReadCloser) (audio.Audio, error) {
	d, err := haj.Decode(r)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(make([]byte, 0, d.Length()))
	_, err = io.Copy(buf, d)
	if err != nil {
		return nil, err
	}
	mformat := mc.Format()
	mformat.SampleRate = uint32(d.SampleRate())
	return audio.EncodeBytes(
		audio.Encoding{
			Data:   buf.Bytes(),
			Format: mformat,
		})
}

// Save will eventually save an audio encoded as an MP3 to r
func (mc *Controller) Save(r io.ReadWriter, a audio.Audio) error {
	return errors.New("Unsupported Functionality")
}

// Format returns the standard mp3 format
func (mc *Controller) Format() audio.Format {
	return format
}
