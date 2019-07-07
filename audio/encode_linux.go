//+build linux

package audio

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/yobert/alsa"
)

type alsaAudio struct {
	*Encoding
	*alsa.Device
	bytesPerPeriod int
	period         int
	playProgress   int
	stopCh         chan struct{}
	playing        bool
	playCh         chan error
}

func (aa *alsaAudio) Play() <-chan error {
	// If currently playing, restart
	if aa.playing {
		aa.playProgress = 0
		return aa.playCh
	}
	aa.playing = true
	aa.playCh = make(chan error)
	go func() {
		for {
			var data []byte
			if len(aa.Encoding.Data)-aa.playProgress <= aa.period {
				data = aa.Encoding.Data[aa.playProgress:]
				if aa.Loop {
					delta := aa.period - (len(aa.Encoding.Data) - aa.playProgress)
					data = append(data, aa.Encoding.Data[:delta]...)
				}
			} else {
				data = aa.Encoding.Data[aa.playProgress : aa.playProgress+aa.period]
			}
			if len(data) != 0 {
				err := aa.Device.Write(data, aa.period)
				if err != nil {
					select {
					case aa.playCh <- err:
					default:
					}
					break
				}
			}
			aa.playProgress += aa.period
			if aa.playProgress > len(aa.Encoding.Data) {
				if aa.Loop {
					aa.playProgress %= len(aa.Encoding.Data)
				} else {
					select {
					case aa.playCh <- nil:
					default:
					}
					break
				}
			}
			select {
			case <-aa.stopCh:
				select {
				case aa.playCh <- nil:
				default:
				}
				break
			default:
			}
		}
		aa.playing = false
		aa.playProgress = 0
	}()
	return aa.playCh
}

func (aa *alsaAudio) Stop() error {
	if aa.playing {
		go func() {
			aa.stopCh <- struct{}{}
		}()
	} else {
		return errors.New("Audio not playing, cannot stop")
	}
	return nil
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
	_, err = handle.NegotiateFormat(format)
	if err != nil {
		return nil, err
	}
	_, err = handle.NegotiateRate(int(enc.SampleRate))
	if err != nil {
		return nil, err
	}
	_, err = handle.NegotiateChannels(int(enc.Channels))
	if err != nil {
		return nil, err
	}
	// Default value at recommendation of library
	period, err := handle.NegotiatePeriodSize(2048)
	if err != nil {
		return nil, err
	}
	_, err = handle.NegotiateBufferSize(4096)
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
		stopCh:         make(chan struct{}),
		bytesPerPeriod: period * (int(enc.Bits) / 8),
	}, nil
}

func openDevice() (*alsa.Device, error) {
	cards, err := alsa.OpenCards()
	if err != nil {
		return nil, err
	}
	for i, c := range cards {
		fmt.Println("Card", c)
		dvcs, err := c.Devices()
		if err != nil {
			alsa.CloseCards([]*alsa.Card{c})
			continue
		}
		for _, d := range dvcs {
			fmt.Println("Device", d)
			if d.Type != alsa.PCM || !d.Play {
				continue
			}
			ech := make(chan error)
			go func() {
				ech <- d.Open()
			}()
			select {
			case <-time.After(300 * time.Millisecond):
				continue
			case err := <-ech:
				if err != nil {
					fmt.Println(err)
					continue
				}
			}
			// We've a found a device we can hypothetically use
			// Close all other cards
			alsa.CloseCards(cards[i+1:])
			return d, nil
		}
		alsa.CloseCards([]*alsa.Card{c})
	}
	return nil, errors.New("No valid device found")
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
