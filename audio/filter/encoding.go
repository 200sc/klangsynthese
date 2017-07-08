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

func LeftPan() Encoding {
	return func(enc supports.Encoding) {
		data := enc.GetData()
		// Right/Left only makes sense for 2 channel
		if *enc.GetChannels() != 2 {
			return
		}
		// Zero out one channel
		swtch := int((*enc.GetBitDepth()) / 8)
		d := *data
		for i := 0; i < len(d); i += (2 * swtch) {
			for j := 0; j < swtch; j++ {
				d[i+j] = byte((int(d[i+j]) + int(d[i+j+swtch])) / 2)
				d[i+j+swtch] = 0
			}
		}
		*data = d
	}
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
		d := *data
		for i := 0; i < len(d); i += (2 * swtch) {
			for j := 0; j < swtch; j++ {
				d[i+j+swtch] = byte((int(d[i+j]) + int(d[i+j+swtch])) / 2)
				d[i+j] = 0
			}
		}
		*data = d
	}
}

// Todo: pans that are not absolute
// problem: information loss
// we need to find which channel has more data to pull from

// Volume will magnify the data by mult, increasing or reducing the volume
// of the output sound. For mult <= 1 this should have no unexpected behavior,
// although for mult ~= 1 it might not have any effect. More importantly for
// mult > 1, values may result in the output data clipping over integer overflows,
// which is presumably not desired behavior.
func Volume(mult float64) Encoding {
	return func(enc supports.Encoding) {
		data := enc.GetData()
		d := *data
		byteDepth := int(*enc.GetBitDepth() / 8)
		for i := 0; i < len(d); i += byteDepth {
			switch byteDepth {
			case 2:
				var v int16
				var shift uint16
				for j := 0; j < byteDepth; j++ {
					v += int16(d[i+j]) << shift
					shift += 8
				}
				v3 := round(float64(v) * mult)
				for j := 0; j < byteDepth; j++ {
					d[i+j] = byte(v3 & 255)
					v3 >>= 8
				}
			default:
				// log unsupported bit depth
				// 2 4 and 8 should also be supported, as int8 int32 and int64
			}
		}
		*data = d
	}
}

func round(f float64) int64 {
	if f < 0 {
		return int64(f - .5)
	}
	return int64(f + .5)
}
