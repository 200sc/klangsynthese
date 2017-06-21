package klangsynthese

// FontAudio is an ease-of-use wrapper around an audio
// with an attached font, so that the audio can be played
// with .Play() but can take in the remotely variable
// font filter options.
//
// Note that it is a concious choice for both Font and
// Audio to have a Filter(...Filter) function, so that when
// a FontAudio is in use the user needs to specify which
// element they want to apply a filter on. The alternative would
// be to have two similarly named functions, and its believed
// that fa.Font.Filter(...) and fa.Audio.Filter(...) is
// more or less equivalent to whatever those names would be.
type FontAudio struct {
	Font
	Audio
}

// NewFontAudio returns a *FontAudio.
// For preparation against API changes, using NewFontAudio over FontAudio{}
// is recommended.
func NewFontAudio(f Font, a Audio) *FontAudio {
	return &FontAudio{f, a}
}

// Play is equivalent to FontAudio.Font.Play(fa.Audio)
func (fa *FontAudio) Play() <-chan error {
	return fa.Font.Play(fa.Audio)
}
