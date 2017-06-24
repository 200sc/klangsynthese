package filter

import "github.com/200sc/klangsynthese/audio"

type Data func(*[]byte)
type SupportsData interface {
	GetData() *[]byte
}

func (df Data) Apply(a audio.Audio) audio.Audio {
	if sd, ok := a.(SupportsData); ok {
		df(sd.GetData())
	}
	return a
}
