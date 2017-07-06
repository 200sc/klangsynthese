package audio

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
}
