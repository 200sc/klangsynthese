package klangsynthese

// Audio represents playable, filterable audio data.
type Audio interface {
	// Play returns a channel that will signal when it finishes playing.
	// Looping audio will never send on this channel!
	// The value sent will always be true.
	Play() <-chan bool
	// Filter will return an audio with some desired filters applied
	Filter(...Filter) Audio
	// Stop will stop an ongoing audio
	Stop() error
	Copy() Audio
}