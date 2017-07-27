package filter

import "github.com/200sc/klangsynthese/audio/filter/supports"

const (
	// Todo: move these to initialization function, offer a few pre-initialized
	// ones
	fftFrameSize = 1024
	oversampling = 32
	step         = fftFrameSize / oversampling
	latency      = fftFrameSize - step
)

// PitchShift modifies filtered audio by the input float, between 0.5 and 2.0,
// each end of the spectrum representing octave down and up respectively
func PitchShift(shiftBy float64) Encoding {
	return func(senc supports.Encoding) {
		data := *senc.GetData()
		bitDepth := *senc.GetBitDepth()
		byteDepth := bitDepth / 8
		sampleRate := *senc.GetSampleRate()
		channels := *senc.GetChannels()
		samplesToProcess := len(data) / int((byteDepth * channels))

		out := make([]byte, len(data))
		stack := make([]float64, fftFrameSize)

		frame := make([]float64, fftFrameSize)
		frameIndex := 0
		// for each channel
		for c := 0; c < int(channels); c++ {
			for i := c; i < len(data); i += int(byteDepth * 2) {
				// Get a frame
				frame[frameIndex] = getFloat64(data, i, byteDepth)
				setInt16_f64(out, i, stack[wrapMod(frameIndex-latency, len(stack))])
				frameIndex++

				// A full frame has been obtained
				if frameIndex == fftFrameSize {
					frameIndex = latency

					// Windowing
					// ...

					// Transform
					// ...

					// Analysis
					// ...

					// Processing
					// ...

					// Synthesis
					// ...

					// Remove negative frequencies
					// ...

					// Inverse transform
					// ...

					// Windowing
					// ...

					// Shift accumulator, shift stack
					// ...
				}
			}
		}
	}
}
