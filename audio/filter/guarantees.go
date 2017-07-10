// Package filter provides various audio filters to be applied to audios through the
// Filter() function
package filter

import (
	"github.com/200sc/klangsynthese/audio"
	"github.com/200sc/klangsynthese/audio/filter/supports"
)

// These declarations guarantee that the filters in this package satisfy the filter interface
var (
	_ audio.Filter = SampleRate(func(*uint32) {})
	_ audio.Filter = Data(func(*[]byte) {})
	_ audio.Filter = Loop(func(*bool) {})
	_ audio.Filter = Encoding(func(supports.Encoding) {})
)
