# klangsynthese
Waveform and Audio Synthesis library in Go

[![GoDoc](https://godoc.org/github.com/200sc/klangsynthese?status.svg)](https://godoc.org/github.com/200sc/klangsynthese)
[![Go Report Card](https://goreportcard.com/badge/github.com/200sc/klangsynthese)](https://goreportcard.com/report/github.com/200sc/klangsynthese)

Klangsynthese right now supports a number of features that will work regardless of OS,
with further support planned for OSX as soon as we get our hands on one to test with. 

## Usage

See test files.

## OS specific features

| OS       | Load | Modify | Save   | Play |
| -------- | ---- | ------ | ------ | ---- |
| Windows  | X    | X      |        |  X   |
| Linux    | X    | X      |        |  X   |
| Darwin   | X    | X      |        |      |

## Quick recipe for testing on Linux

This recipe should run the wav test on Linux:

    go get github.com/200sc/klangsynthese
    go get github.com/stretchr/testify/require
    go test github.com/200sc/klangsynthese/wav

## Other features

- [x] Wav support
- [x] Mp3 support
- [x] Flac support
- [ ] Ogg support
- [x] Creating waveforms (Sin, Square, Saw, ...)
- [x] Filtering audio samples
- [x] Creating Sequences of audio samples to play music
- [x] Importable Bosca Ceoil files (.ceol)
- [x] Importable DLS files 
- [ ] Importable arbitrary instruments
- [ ] Other Importable soundfonts (.sf2...)
