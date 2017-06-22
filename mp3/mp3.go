package mp3

import (
	"bytes"
	"io"

	"github.com/200sc/klangsynthese/audio"

	haj "github.com/hajimehoshi/go-mp3"
)

var format = audio.Format{
	SampleRate: 44100,
	Bits:       16,
	Channels:   2,
}

// This should store device information?
type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

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
	format := mc.Format()
	format.SampleRate = uint32(d.SampleRate())
	return audio.EncodeBytes(
		audio.Encoding{
			Data:   buf.Bytes(),
			Format: format,
		})
}

func (mc *Controller) Save(r io.ReadWriter, a audio.Audio) error {
	return nil
}

func (mc *Controller) Format() audio.Format {
	return format
}
