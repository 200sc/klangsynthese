package klangsynthese

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
