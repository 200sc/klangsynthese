package audio

// A Filter takes an input audio and returns some new Audio from them.
// This usage implies that Audios can be copied, and that Audios have
// available information to be generically modified by a Filter. The
// functions for these capabilities are yet fleshed out. It's worth
// considering whether a Filter modifies in place. The answer is
// probably yes:
// a.Filter(fs) would modify a in place
// a.Copy().Filter(fs) would return a new audio
// Specific audio implementations could not follow this, however.
type Filter func(Audio) Audio

func Loop(a Audio) Audio {
	enc := a.GetEncoding()
	enc.loop = true
	return a
}

// NoLoop is only meaningful on audio that has already had Loop filtered on it
func NoLoop(a Audio) Audio {
	enc := a.GetEncoding()
	enc.loop = false
	return a
}
