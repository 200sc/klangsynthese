package audio

// Encoding contains all information required to convert raw data
// (currently assumed PCM data but that may/will change) into playable Audio
type Encoding struct {
	// Consider: non []byte data?
	Data []byte
	Format
}

// Encoding returns itself
func (enc *Encoding) GetEncoding() Encoding {
	return *enc
}

// We need access to everything we used to create the buffer in order to copy it
func (enc *Encoding) Copy() Audio {
	// The error is currently ignored (because presumably you have
	// already created the Audio and are copying it) but that may
	// change in the future (the reason it would not is to keep the
	// api easy, it's troublesome to have to copy on a separate line)
	a, _ := EncodeBytes(*enc)
	return a
}

// Format stores the variables which are presumably
// constant for any given type of audio (wav / mp3 / flac ...)
type Format struct {
	SampleRate uint32
	Channels   uint16
	Bits       uint16
}
