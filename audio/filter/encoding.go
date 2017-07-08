package filter

import (
	"github.com/200sc/klangsynthese/audio"
	"github.com/200sc/klangsynthese/audio/filter/supports"
)

// Encoding filters are functions on any combination of the values
// in an audio.Encoding
type Encoding func(supports.Encoding)

// Apply checks that the given audio supports Encoding, filters if it
// can, then returns
func (enc Encoding) Apply(a audio.Audio) (audio.Audio, error) {
	if senc, ok := a.(supports.Encoding); ok {
		enc(senc)
		return a, nil
	}
	return a, supports.NewUnsupported([]string{"Encoding"})
}

func RightPan() Encoding {
	return func(enc supports.Encoding) {
		data := enc.GetData()
		// Right/Left only makes sense for 2 channel
		if *enc.GetChannels() != 2 {
			return
		}
		// Zero out one channel
		swtch := int((*enc.GetBitDepth()) / 8)
		zero := false
		d := *data
		for i := range d {
			if i%swtch == 0 {
				zero = !zero
			}
			if zero {
				d[i] = 0
			}
		}
		*data = d
	}
}

func LeftPan() Encoding {
	return func(enc supports.Encoding) {
		data := enc.GetData()
		// Right/Left only makes sense for 2 channel
		if *enc.GetChannels() != 2 {
			return
		}
		// Zero out one channel
		swtch := int((*enc.GetBitDepth()) / 8)
		zero := true
		d := *data
		for i := range d {
			if i%swtch == 0 {
				zero = !zero
			}
			if zero {
				d[i] = 0
			}
		}
		*data = d
	}
}

// Todo: pans that are not absolute
// problem: information loss
// we need to find which channel has more data to pull from
