package mp3

import (
	"bytes"
	"fmt"
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
	read, err := io.Copy(buf, d)
	b := buf.Bytes()
	fmt.Println(buf.Len())
	fmt.Println(read, err)
	format := mc.Format()
	format.SampleRate = uint32(d.SampleRate())
	fmt.Println(d)
	fmt.Println("Mp3 format:", format)
	return audio.EncodeBytes(
		audio.Encoding{
			Data:   b,
			Format: format,
		})
}

func (mc *Controller) Save(r io.ReadWriter, a audio.Audio) error {
	return nil
}

func (mc *Controller) Format() audio.Format {
	return format
}
