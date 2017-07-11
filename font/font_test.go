package font

import (
	"testing"
	"time"

	"github.com/200sc/klangsynthese/audio"
	"github.com/200sc/klangsynthese/audio/filter"
	"github.com/200sc/klangsynthese/synth"
	"github.com/200sc/klangsynthese/wav"
	"github.com/stretchr/testify/assert"
)

func TestFont1(t *testing.T) {
	f := New().Filter(filter.Volume(.25))
	a, err := wav.NewController().Wave(synth.Sin(synth.A4, 2, 2000))
	assert.Nil(t, err)
	fa := NewAudio(f, a.(audio.FullAudio))
	fa.Play()
	fa.Font.Filter(filter.Volume(.5))
	time.Sleep(2 * time.Second)
	fa.Play()
	fa.Font.Filter(filter.Volume(.75))
	time.Sleep(2 * time.Second)
	fa.Play()
	time.Sleep(2 * time.Second)
}
