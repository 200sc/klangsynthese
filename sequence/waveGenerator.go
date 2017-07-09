package sequence

import (
	"time"

	"github.com/200sc/klangsynthese/audio"
	"github.com/200sc/klangsynthese/synth"
	"github.com/200sc/klangsynthese/wav"
)

type WaveGenerator struct {
	ChordPattern
	PitchPattern
	WavePattern
	VolumePattern
	HoldPattern
	Length
	Tick
	Loop
}

func NewWaveGenerator(opts ...Option) *WaveGenerator {
	wg := &WaveGenerator{}
	for _, opt := range opts {
		opt(wg)
	}
	return wg
}

func (wg *WaveGenerator) Generate() *Sequence {
	sq := &Sequence{}
	sq.Ticker = time.NewTicker(time.Duration(wg.Tick))
	sq.tickDuration = time.Duration(wg.Tick)
	sq.loop = bool(wg.Loop)
	sq.stopCh = make(chan error)
	if wg.Length == 0 {
		if len(wg.PitchPattern) != 0 {
			wg.Length = Length(len(wg.PitchPattern))
		} else if len(wg.ChordPattern.Pitches) != 0 {
			wg.Length = Length(len(wg.ChordPattern.Pitches))
		}
		// else whoops, there's no length
	}
	if len(wg.HoldPattern) == 0 {
		wg.HoldPattern = []time.Duration{sq.tickDuration}
	}
	sq.Pattern = make([]audio.Multi, wg.Length)

	controller := wav.NewController()

	volumeIndex := 0
	waveIndex := 0
	if len(wg.PitchPattern) != 0 {
		pitchIndex := 0
		holdIndex := 0
		for i := range sq.Pattern {
			p := wg.PitchPattern[pitchIndex]
			if p != synth.Rest {
				a, _ := controller.Wave(
					wg.WavePattern[waveIndex](
						p,
						wg.HoldPattern[holdIndex].Seconds(),
						wg.VolumePattern[volumeIndex],
					))
				sq.Pattern[i] = audio.Multi{[]audio.Audio{a}}
			} else {
				sq.Pattern[i] = audio.Multi{[]audio.Audio{}}
			}
			pitchIndex = (pitchIndex + 1) % len(wg.PitchPattern)
			volumeIndex = (volumeIndex + 1) % len(wg.VolumePattern)
			waveIndex = (waveIndex + 1) % len(wg.WavePattern)
			holdIndex = (holdIndex + 1) % len(wg.HoldPattern)
		}
	} else if len(wg.ChordPattern.Pitches) != 0 {
		chordIndex := 0
		for i := range sq.Pattern {
			mult := audio.Multi{}
			for j, p := range wg.ChordPattern.Pitches[chordIndex] {
				a, _ := controller.Wave(
					wg.WavePattern[waveIndex](
						p,
						wg.ChordPattern.Holds[chordIndex][j].Seconds(),
						wg.VolumePattern[volumeIndex],
					))
				mult.Audios = append(mult.Audios, a)
			}
			sq.Pattern[i] = mult
			waveIndex = (waveIndex + 1) % len(wg.WavePattern)
			volumeIndex = (volumeIndex + 1) % len(wg.VolumePattern)
			chordIndex = (chordIndex + 1) % len(wg.ChordPattern.Pitches)
		}
	}
	return sq
}
