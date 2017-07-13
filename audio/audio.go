// Package audio provides audio playing and encoding support
package audio

import (
	"time"

	"github.com/200sc/klangsynthese/audio/filter/supports"
)

// Audio represents playable, filterable audio data.
type Audio interface {
	// Play returns a channel that will signal when it finishes playing.
	// Looping audio will never send on this channel!
	// The value sent will always be true.
	Play() <-chan error
	// Filter will return an audio with some desired filters applied
	Filter(...Filter) (Audio, error)
	MustFilter(...Filter) Audio
	// Stop will stop an ongoing audio
	Stop() error

	// Implementing struct-- encoding
	Copy() (Audio, error)
	MustCopy() Audio
	PlayLength() time.Duration
}

// FullAudio supports all the built in filters
type FullAudio interface {
	Audio
	supports.Encoding
	supports.Loop
}
