package sequence

// A Generator stores settings to create a sequence
type Generator interface {
	Generate() *Sequence
}
