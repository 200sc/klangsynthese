package audio

// Format stores the variables which are presumably
// constant for any given type of audio (wav / mp3 / flac ...)
type Format struct {
	SampleRate uint32
	Channels   uint16
	Bits       uint16
}

// GetSampleRate satisfies supports.SampleRate
func (f *Format) GetSampleRate() *uint32 {
	return &f.SampleRate
}

func (f *Format) GetChannels() *uint16 {
	return &f.Channels
}

func (f *Format) GetBitDepth() *uint16 {
	return &f.Bits
}
