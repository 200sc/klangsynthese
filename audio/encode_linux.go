//+build linux

package audio

import (
	"sync"

	"github.com/pkg/errors"
	"github.com/yobert/alsa"
)

type alsaAudio struct {
	*Encoding
	*alsa.Device
	bytesPerPeriod int
	period         int
	playProgress   int
	playing        bool
	playMutex      sync.Mutex
}

func (aa *alsaAudio) Play() <-chan error {
	ch := make(chan error)
	// If currently playing, restart
	aa.playMutex.Lock()
	if aa.playing {
		aa.playProgress = 0
		aa.playMutex.Unlock()
		return
	}
	aa.playMutex.Unlock()
	go func() {
		for {
			var data []byte
			if len(aa.Encoding.Data)-aa.playProgress >= aa.period {
				data = aa.Encoding.Data[aa.playProgress:]
				if aa.Loop {
					delta := aa.period - (len(aa.Encoding.Data)-aa.playProgress)
					data = append(data, aa.Encoding.Data[:delta)
				}
			} else {
				data = aa.Encoding.Data[aa.playProgress : aa.playProgress+aa.period]
			}
			err := aa.Device.Write(data, aa.period)
			if err != nil {
				ch <- err
				return
			}
			// Consider: its racy, but we could remove this lock and risk
			// skipping the first period of the audio on concurrent play requests
			aa.playMutex.Lock()
			aa.playProgress += aa.period
			if aa.playProgress > len(aa.Encoding.Data) {
				if aa.Loop {
					aa.playProgress %= len(aa.Encoding.Data)
				} else {
					aa.playMutex.Unlock()
					return 
				}
			}
			aa.playMutex.Unlock()
		}
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
			if consErr == nil {
				consErr = err
			} else {
				consErr = errors.New(err.Error() + ":" + consErr.Error())
			}
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
	handle, err := openDevice()
	if err != nil {
		return nil, err
	}
	//err := handle.Open("default", alsa.StreamTypePlayback, alsa.ModeBlock)
	//if err != nil {
	//	return nil, err
	//}
	// Todo: annotate these errors with more info
	format, err := alsaFormat(enc.Bits)
	if err != nil {
		return nil, err
	}
	_, err := handle.NegotiateFormat(format)
	if err != nil {
		return nil, err
	}
	_, err := handle.NegotiateRate(int(enc.SampleRate))
	if err != nil {
		return nil, err
	}
	_, err := handle.NegotiateChannels(int(enc.Channels))
	if err != nil {
		return nil, err
	}
	// Default value at recommendation of library
	period, err := handle.NegotiatePeriodSize(2048)
	if err != nil {
		return nil, err
	}
	_, err := handle.NegotiateBufferSize(4096)
	if err != nil {
		return nil, err
	}
	err = handle.Prepare()
	if err != nil {
		return nil, err
	}
	return &alsaAudio{
		Encoding:       &enc,
		Device:         handle,
		period:         period,
		bytesPerPeriod: period * (enc.Bits / 8),
	}, nil
}

func openDevice() (*alsa.Device, error) {
	cards, err := alsa.OpenCards()
	if err != nil {
		return nil, err
	}
	for i, c := range cards {
		dvcs, err := c.Devices()
		if err != nil {
			alsa.CloseCards([]*alsa.Card{c})
			continue
		}
		for j, d := range dvcs {
			if d.Type != alsa.PCM || !d.Play {
				d.Close()
				continue
			}
			// We've a found a device we can hypothetically use
			// Close all other cards and devices
			for h := j + 1; h < len(dvcs); h++ {
				dvcs[h].Close()
			}
			alsa.CloseCards(cards[i+1:])
			return d, d.Open()
		}
		alsa.CloseCards([]*alsa.Card{c})
	}
}

func alsaFormat(bits uint16) (alsa.FormatType, error) {
	switch bits {
	case 8:
		return alsa.S8, nil
	case 16:
		return alsa.S16_LE, nil
	}
	return 0, errors.New("Undefined alsa format for encoding bits")
}
