package sequence

import (
	"time"

	"github.com/200sc/klangsynthese/audio"
	"github.com/200sc/klangsynthese/synth"
	"github.com/200sc/klangsynthese/wav"
)

type WaveGenerator struct {
	Fn            synth.Wave
	Tick          time.Duration
	PitchPattern  []synth.Pitch
	VolumePattern []synth.Volume
	Loop          bool
}

func (wg *WaveGenerator) Generate() *Sequence {
	sq := &Sequence{}
	sq.Ticker = time.NewTicker(wg.Tick)
	sq.tickDuration = wg.Tick
	sq.loop = wg.Loop
	sq.stopCh = make(chan error)
	patternLength := len(wg.PitchPattern)
	if len(wg.VolumePattern) > patternLength {
		patternLength = len(wg.VolumePattern)
	}
	sq.Pattern = make([]audio.Multi, patternLength)

	controller := wav.NewController()

	pitchIndex := 0
	volumeIndex := 0
	tickSeconds := wg.Tick.Seconds()
	for i := range sq.Pattern {
		p := wg.PitchPattern[pitchIndex]
		if p != synth.Rest {
			a, _ := controller.Wave(wg.Fn(
				p,
				tickSeconds,
				wg.VolumePattern[volumeIndex],
			))
			sq.Pattern[i] = audio.Multi{[]audio.Audio{a}}
		} else {
			sq.Pattern[i] = audio.Multi{[]audio.Audio{}}
		}
		pitchIndex = (pitchIndex + 1) % len(wg.PitchPattern)
		volumeIndex = (volumeIndex + 1) % len(wg.VolumePattern)
	}
	return sq
}
