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
