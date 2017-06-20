package klangsynthese

// A Synth returns raw data for a controller to process into audio
// Consider: should this be named SynthFn? Wave? WaveFn?
type Synth func() []byte
