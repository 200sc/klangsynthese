package filter

import (
	"github.com/200sc/klangsynthese/audio"
	"github.com/200sc/klangsynthese/audio/filter/supports"
)

// Data filters are functions on []byte types
type Data func(*[]byte)

// Apply checks that the given audio supports Data, filters if it
// can, then returns
func (df Data) Apply(a audio.Audio) (audio.Audio, error) {
	if sd, ok := a.(supports.Data); ok {
		df(sd.GetData())
		return a, nil
	}
	return a, supports.NewUnsupported([]string{"Data"})
}

// Volume will magnify the data by mult, increasing or reducing the volume
// of the output sound. For mult <= 1 this should have no unexpected behavior,
// although for mult ~= 1 it might not have any effect. More importantly for
// mult > 1, values may result in the output data clipping over integer overflows,
// which is presumably not desired behavior.
func Volume(mult float64) Data {
	return func(bp *[]byte) {
		b := *bp
		for i, v := range b {
			b[i] = byte(float64(int8(v)) * mult)
		}
		*bp = b
	}
}
