package filter

import (
	"fmt"

	"github.com/200sc/klangsynthese/audio"
	"github.com/200sc/klangsynthese/audio/filter/supports"
)

// Data filters are functions on []byte types
type Data func(*[]byte)

// Apply checks that the given audio supports Data, filters if it
// can, then returns
func (df Data) Apply(a audio.Audio) (audio.Audio, error) {
	if sd, ok := a.(supports.Data); ok {
		fmt.Println((*sd.GetData())[2000])
		df(sd.GetData())
		fmt.Println((*sd.GetData())[2000])
		return a, nil
	}
	return a, supports.NewUnsupported([]string{"Data"})
}

func Volume(mult float64) Data {
	return func(bp *[]byte) {
		b := *bp
		for i, v := range b {
			b[i] = byte(float64(int8(v)) * mult)
		}
		*bp = b
	}
}
