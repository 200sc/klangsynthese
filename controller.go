package klangsynthese

import "io"

// Controller represents the ability to generate Audio from data, or from
// synthesis options
type Controller interface {
	Wave(Synth) (Audio, error)
	Load(io.Reader) (Audio, error)
	Save(io.Writer, Audio) error
}
