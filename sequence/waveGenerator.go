package sequence

import (
	"fmt"
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
	fmt.Println(tickSeconds, "Seconds")
	for i := range sq.Pattern {
		a, _ := controller.Wave(wg.Fn(
			wg.PitchPattern[pitchIndex],
			tickSeconds,
			wg.VolumePattern[volumeIndex],
		))
		sq.Pattern[i] = audio.Multi{[]audio.Audio{a}}
		// if err != nil {
		// return err
		// }
		pitchIndex = (pitchIndex + 1) % len(wg.PitchPattern)
		volumeIndex = (volumeIndex + 1) % len(wg.VolumePattern)
	}
	return sq
}
