//+build linux

package audio

import (
	"github.com/pkg/errors"
	"github.com/tryphon/alsa-go"
)

type alsaAudio struct {
	*Encoding
	*alsa.Handle
}

func (aa *alsaAudio) Play() <-chan error {
	ch := make(chan error)
	go func() {
		// Todo: loop? library does not export loop
		_, err := aa.Handle.Write(aa.Encoding.Data)
		ch <- err
	}()
	return ch
}

func (aa *alsaAudio) Stop() error {
	// Todo: don't just pause man, actually stop
	// library we are using does not export stop
	return aa.Pause()
}

func (aa *alsaAudio) Filter(fs ...Filter) (Audio, error) {
	var a Audio = aa
	var err, consErr error
	for _, f := range fs {
		a, err = f.Apply(a)
		if err != nil {
			consErr = errors.New(err.Error() + ":" + consErr.Error())
		}
	}
	return aa, consErr
}

// MustFilter acts like Filter, but ignores errors (it does not panic,
// as filter errors are expected to be non-fatal)
func (aa *alsaAudio) MustFilter(fs ...Filter) Audio {
	a, _ := aa.Filter(fs...)
	return a
}

func EncodeBytes(enc Encoding) (Audio, error) {
	handle := alsa.New()
	err := handle.Open("default", alsa.StreamTypePlayback, alsa.ModeBlock)
	if err != nil {
		return nil, err
	}

	handle.SampleFormat = alsaFormat(enc.Bits)
	handle.SampleRate = int(enc.SampleRate)
	handle.Channels = int(enc.Channels)
	err = handle.ApplyHwParams()
	if err != nil {
		return nil, err
	}
	return &alsaAudio{
		&enc,
		handle,
	}, nil
}

func alsaFormat(bits uint16) alsa.SampleFormat {
	switch bits {
	case 8:
		return alsa.SampleFormatS8
	case 16:
		return alsa.SampleFormatS16LE
	}
	return alsa.SampleFormatUnknown
}
