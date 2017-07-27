# klangsynthese
Waveform and Audio Synthesis library in Go

[![GoDoc](https://godoc.org/github.com/200sc/klangsynthese?status.svg)](https://godoc.org/github.com/200sc/klangsynthese)
[![Go Report Card](https://goreportcard.com/badge/github.com/200sc/klangsynthese)](https://goreportcard.com/report/github.com/200sc/klangsynthese)

Klangsynthese right now supports a number of features that will work regardless of OS,
and a number of features specific to Windows where the hope is to move support to Linux
and Darwin.

## Usage

See test files.

## OS specific features

| OS       | Load | Modify | Save   | Play |
| -------- | ---- | ------ | ------ | ---- |
| Windows  | X    | X      |        |  X   |
| Linux    | X    | X      |        |  ?   |
| Darwin   | X    | X      |        |      |

To develop with linux you'll need alsa:

`sudo apt-get install alsa-base libasound2-dev`

Binaries built with this will probably need alsa-base as well to run on Linux.

## Other features

- [x] Wav support
- [x] Mp3 support
- [x] Flac support?
- [ ] Ogg support
- [x] Creating waveforms (Sin, Square, Saw, ...)
- [x] Filtering audio samples
- [x] Creating Sequences of audio samples to play music
- [x] Importable Bosca Ceoil files (.ceol)
- [x] Importable DLS files 
- [ ] Importable arbitrary instruments
- [ ] Other Importable soundfonts (.sf2...)
