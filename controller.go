package klangsynthese

import (
	"io"

	"github.com/200sc/klangsynthese/audio"
)

// Controller represents the ability to generate Audio from data, or from
// synthesis options
type Controller interface {
	Wave(Synth) (audio.Audio, error)
	Load(io.ReadCloser) (audio.Audio, error)
	Save(io.WriteCloser, audio.Audio) error
	Format() audio.Format
}
