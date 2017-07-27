package filter

import (
	"fmt"
	"math"

	"github.com/200sc/klangsynthese/audio/filter/supports"
)

const (
	// Todo: move these to initialization function, offer a few pre-initialized
	// ones
	fftFrameSize = 1024
	oversampling = 32
	step         = fftFrameSize / oversampling
	latency      = fftFrameSize - step
)

/*****************************************************************************
* HOME URL: http://blogs.zynaptiq.com/bernsee
* KNOWN BUGS: none
*
* SYNOPSIS: Routine for doing pitch shifting while maintaining
* duration using the Short Time Fourier Transform.
*
* DESCRIPTION: The routine takes a pitchShift factor value which is between 0.5
* (one octave down) and 2. (one octave up). A value of exactly 1 does not change
* the pitch. numSampsToProcess tells the routine how many samples in indata[0...
* numSampsToProcess-1] should be pitch shifted and moved to outdata[0 ...
* numSampsToProcess-1]. The two buffers can be identical (ie. it can process the
* data in-place). fftFrameSize defines the FFT frame size used for the
* processing. Typical values are 1024, 2048 and 4096. It may be any value <=
* MAX_FRAME_LENGTH but it MUST be a power of 2. osamp is the STFT
* oversampling factor which also determines the overlap between adjacent STFT
* frames. It should at least be 4 for moderate scaling ratios. A value of 32 is
* recommended for best quality. sampleRate takes the sample rate for the signal
* in unit Hz, ie. 44100 for 44.1 kHz audio. The data passed to the routine in
* indata[] should be in the range [-1.0, 1.0), which is also the output range
* for the data, make sure you scale the data accordingly (for 16bit signed integers
* you would have to divide (and multiply) by 32768).
*
* COPYRIGHT 1999-2015 Stephan M. Bernsee <s.bernsee [AT] zynaptiq [DOT] com>
*
* 						The Wide Open License (WOL)
*
* Permission to use, copy, modify, distribute and sell this software and its
* documentation for any purpose is hereby granted without fee, provided that
* the above copyright notice and this license appear in all source copies.
* THIS SOFTWARE IS PROVIDED "AS IS" WITHOUT EXPRESS OR IMPLIED WARRANTY OF
* ANY KIND. See http://www.dspguru.com/wol.htm for more information.
*
*****************************************************************************/
// As is standard with translations of this code to other languages,
// Go translation copyright Patrick Stephen 2017

// PitchShift modifies filtered audio by the input float, between 0.5 and 2.0,
// each end of the spectrum representing octave down and up respectively
func PitchShift(shiftBy float64) Encoding {
	return func(senc supports.Encoding) {
		data := *senc.GetData()
		bitDepth := *senc.GetBitDepth()
		byteDepth := bitDepth / 8
		sampleRate := *senc.GetSampleRate()
		channels := *senc.GetChannels()

		// Jeeez
		out := make([]byte, len(data))
		copy(out, data)
		stack := make([]float64, fftFrameSize)
		workBuffer := make([]float64, 2*fftFrameSize)
		magnitudes := make([]float64, fftFrameSize)
		frequencies := make([]float64, fftFrameSize)
		synthMagnitudes := make([]float64, fftFrameSize)
		synthFrequencies := make([]float64, fftFrameSize)
		lastPhase := make([]float64, fftFrameSize/2+1)
		sumPhase := make([]float64, fftFrameSize/2+1)
		outAcc := make([]float64, 2*fftFrameSize)

		freqPerBin := float64(sampleRate) / float64(fftFrameSize)
		expected := 2 * math.Pi * float64(step) / float64(fftFrameSize)

		window := make([]float64, fftFrameSize)
		windowFactors := make([]float64, fftFrameSize)
		t := 0.0
		for i := 0; i < fftFrameSize; i++ {
			w := -0.5*math.Cos(t) + .5
			window[i] = w
			windowFactors[i] = w * (2.0 / (fftFrameSize * oversampling))
			t += (math.Pi * 2) / fftFrameSize
		}

		frame := make([]float64, fftFrameSize)
		frameIndex := latency

		// End jeeeez

		// for each channel individually
		for c := 0; c < int(channels); c++ {
			// convert this to a channel-specific float64 buffer
			f64in := make([]float64, len(data)/int(channels))
			f64out := f64in
			for i := c * int(byteDepth); i < len(data); i += int(byteDepth * 2) {
				f64in[i/int(byteDepth*2)] = getFloat64(data, i, byteDepth)
			}
			// At this point, we are confident we are correct
			for i := 0; i < len(f64in); i++ {
				// Get a frame
				frame[frameIndex] = f64in[i]
				f64out[i] = stack[frameIndex-latency]
				frameIndex++

				// A full frame has been obtained
				if frameIndex >= fftFrameSize {
					frameIndex = latency

					// Windowing
					for k := 0; k < fftFrameSize; k++ {
						workBuffer[2*k] = frame[k] * window[k]
						workBuffer[(2*k)+1] = 0
					}

					ShortTimeFourierTransform(workBuffer, fftFrameSize, -1)

					// Analysis
					for k := 0; k <= fftFrameSize/2; k++ {
						real := workBuffer[2*k]
						imag := workBuffer[(2*k)+1]

						magn := 2 * math.Sqrt(real*real+imag*imag)
						magnitudes[k] = magn

						phase := math.Atan2(imag, real)

						diff := phase - lastPhase[k]
						lastPhase[k] = phase

						diff -= float64(k) * expected

						deltaPhase := int(diff * (1 / math.Pi))
						if deltaPhase >= 0 {
							deltaPhase += deltaPhase & 1
						} else {
							deltaPhase -= deltaPhase & 1
						}

						diff -= math.Pi * float64(deltaPhase)
						diff *= oversampling / (math.Pi * 2)
						diff = (float64(k) + diff) * freqPerBin

						frequencies[k] = diff
					}

					// Processing
					for k := 0; k < fftFrameSize; k++ {
						synthMagnitudes[k] = 0
						synthFrequencies[k] = 0
					}

					for k := 0; k < fftFrameSize/2; k++ {
						l := int(float64(k) * shiftBy)
						if l < fftFrameSize/2 {
							synthMagnitudes[l] += magnitudes[k]
							synthFrequencies[l] = frequencies[k] * shiftBy
						}
					}

					// Synthesis
					for k := 0; k <= fftFrameSize/2; k++ {
						magn := synthMagnitudes[k]
						tmp := synthFrequencies[k]
						tmp -= float64(k) * freqPerBin
						tmp /= freqPerBin
						tmp *= 2 * math.Pi / oversampling
						tmp += float64(k) * expected
						sumPhase[k] += tmp

						workBuffer[2*k] = magn * math.Cos(sumPhase[k])
						workBuffer[(2*k)+1] = magn * math.Sin(sumPhase[k])
					}

					// Remove negative frequencies
					// I don't get how we know these ones are negative
					// also this looks like it's going to overflow the slice
					for k := fftFrameSize + 2; k < 2*fftFrameSize; k++ {
						workBuffer[k] = 0.0
					}

					ShortTimeFourierTransform(workBuffer, fftFrameSize, 1)

					// Windowing
					for k := 0; k < fftFrameSize; k++ {
						outAcc[k] += windowFactors[k] * workBuffer[2*k]
					}
					for k := 0; k < step; k++ {
						stack[k] = outAcc[k]
					}

					// Shift accumulator, shift frame
					for k := 0; k < fftFrameSize; k++ {
						outAcc[k] = outAcc[k+step]
					}

					for k := 0; k < latency; k++ {
						frame[k] = frame[k+step]
					}
				}
			}
			// remap this f64in to the output
			for i := c * int(byteDepth); i < len(data); i += int(byteDepth * 2) {
				setInt16_f64(out, i, f64in[i/int(byteDepth*2)])
			}
		}
		datap := senc.GetData()
		*datap = out
		for i := 0; i < 80; i += 2 {
			fmt.Print(getInt16(*datap, i))
			fmt.Print(" ")
		}
	}
}

func ShortTimeFourierTransform(data []float64, fftFrameSize, sign int) {
	for i := 2; i < 2*(fftFrameSize-2); i += 2 {
		j := 0
		for bitm := 2; bitm < 2*fftFrameSize; bitm <<= 1 {
			if (i & bitm) != 0 {
				j++
			}
			j <<= 1
		}
		if i < j {
			data[j], data[i] = data[i], data[j]
			data[j+1], data[i+1] = data[i+1], data[j+1]
		}
	}
	max := int(math.Log(float64(fftFrameSize))/math.Log(2) + .5)
	le := 2
	for k := 0; k < max; k++ {
		le <<= 1
		le2 := le >> 1
		ur := 1.0
		ui := 0.0
		arg := math.Pi / float64(le2>>1)
		wr := math.Cos(arg)
		wi := float64(sign) * math.Sin(arg)
		for j := 0; j < le2; j += 2 {
			for i := j; i < 2*fftFrameSize; i += le {
				//fmt.Println("k,j,i,le,le2", k, j, i, le, le2, max)
				tr := data[i+le2]*ur - data[i+le2+1]*ui
				ti := data[i+le2]*ui + data[i+le2+1]*ur
				data[i+le2] = data[i] - tr
				data[i+le2+1] = data[i+1] - ti
				data[i] += tr
				data[i+1] += ti
			}
			tmp := ur*wr - ui*wi
			ui = ur*wi + ui*wr
			ur = tmp
		}
	}
}
