# klangsynthese
Waveform and Audio Synthesis library in Go

Klangsynthese right now supports a number of features that will work regardless of OS,
and a number of features specific to Windows where the hope is to move support to Linux
and Darwin.

## OS specific features

| OS       | Wav        | MP3       | FLAC   | OGG |
| -------- | ---------- | --------- | ------ | --- |
| Windows  | Load+Play  | Load+Play |        |     |
| Linux    | Load       | Load      |        |     |
| Darwin   | Load       | Load      |        |     |

This library wants to be a zero-dependency library (besides Go), which causes issues for
Linux, and that is why there is no Linux support yet. What will likely happen with this library 
is that we will have a sad, temporary dependencied solution for non-Windows that we will
eventually replace with a custom ALSA audio driver. 

## Other features

- [x] Creating waveforms (Sin, Square, Saw, ...)
- [x] Filtering audio samples
- [x] Creating Sequences of audio samples to play music
- [x] Importable Bosca Ceoil files (.ceol)
- [ ] Importable arbitrary instruments
- [ ] Other Importable soundfonts (.sf2, .dls ...)
