package synth

// Volume is a helper type to avoid maximum / minimum volume limits
type Volume uint8

// Volume const
const (
	MinVolume Volume = 1
	MaxVolume Volume = 127
)
