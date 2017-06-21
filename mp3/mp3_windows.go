// +build windows

package mp3

import (
	"io"

	"github.com/200sc/klangsynthese/audio"

	haj "github.com/hajimehoshi/go-mp3"
)

type Controller struct {
}

func (mc *Controller) Load(r io.ReadCloser) (audio.Audio, error) {
	d, err := haj.Decode(r)
	if err != nil {
		return nil, err
	}
	b := make([]byte, d.Length())
	_, err = d.Read(b)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (mc *Controller) Save(r io.ReadWriter, a klang.Audio) error {
	return nil
}
