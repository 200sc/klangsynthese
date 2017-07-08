package supports

// Data types support filters that manipulate their raw audio data
type Data interface {
	GetData() *[]byte
}

// Loop types support filters that manipulate whether they loop
type Loop interface {
	GetLoop() *bool
}

// SampleRate types support filters that manipulate their SampleRate
type SampleRate interface {
	GetSampleRate() *uint32
}

// Pan types support panning filters for playing in stereo
type Pan interface {
	GetPan() float64
	SetPan(float64)
}
