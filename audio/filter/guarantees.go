package filter

import "github.com/200sc/klangsynthese/audio"

// These declarations guarantee that the filters in this package satisfy the filter interface
var (
	_ audio.Filter = SampleRate(func(*uint32) {})
	_ audio.Filter = Data(func(*[]byte) {})
	_ audio.Filter = Loop(func(*bool) {})
)
