package synth

// Volume is a helper type to avoid maximum / minimum volume limits
type Volume uint16

// Volume const
const (
	MinVolume Volume = 1
	MaxVolume Volume = (65536 / 2) - 1
)
