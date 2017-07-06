package audio

// Encoding contains all information required to convert raw data
// (currently assumed PCM data but that may/will change) into playable Audio
type Encoding struct {
	// Consider: non []byte data?
	// Consider: should Data be a type just like Format and CanLoop?
	Data []byte
	Format
	CanLoop
}

// HasEncoding is the type of any audio with an explicit encoding
// i.e. individual audio samples as opposed to sequences or composites.
type HasEncoding interface {
	GetEncoding() *Encoding
}

// GetEncoding returns itself
func (enc *Encoding) GetEncoding() *Encoding {
	return enc
}

// Copy returns an audio encoded from this encoding.
// Consider: Copy might be tied to HasEncoding
func (enc *Encoding) Copy() (Audio, error) {
	return EncodeBytes(*enc)
}

// MustCopy acts like Copy, but will panic if err != nil
func (enc *Encoding) MustCopy() Audio {
	a, err := EncodeBytes(*enc)
	if err != nil {
		panic(err)
	}
	return a
}

// GetData satisfies filter.SupportsData
func (enc *Encoding) GetData() *[]byte {
	return &enc.Data
}
