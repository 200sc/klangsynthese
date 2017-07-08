package synth

import (
	"testing"
	"time"

	"github.com/200sc/klangsynthese/audio/filter"
	"github.com/200sc/klangsynthese/wav"
	"github.com/stretchr/testify/assert"
)

func TestSinWav(t *testing.T) {
	testWave(t, Sin(A4, 2, 2000))
}

func TestSquareWav(t *testing.T) {
	testWave(t, Square(A4, 2, 2000))
}

func TestSawWav(t *testing.T) {
	testWave(t, Saw(A4, 2, 2000))
}

func TestTriangleWav(t *testing.T) {
	testWave(t, Triangle(A4, 2, 2000))
}

func TestPulseWav(t *testing.T) {
	testWave(t, Pulse(8)(A4, 2, 2000))
}

func TestAdd(t *testing.T) {
	testWave(t,
		//	i.e harmonics
		Add(Sin(A4, 2, 16),
			Sin(A4*2, 2, 16),
			Sin(A4*3, 2, 16),
			Sin(A4*4, 2, 16),
			Sin(A4*5, 2, 16),
			Sin(A4*6, 2, 16),
		))
}

func TestVolume(t *testing.T) {
	a, _ := wav.NewController().Wave(Sin(A4, 1, 2000))
	a2, err := a.MustCopy().Filter(filter.Volume(.25))
	a3, _ := a.MustCopy().Filter(filter.VolumeRight(.5))
	a4, _ := a.MustCopy().Filter(filter.VolumeLeft(.5))
	assert.Nil(t, err)
	a.Play()
	time.Sleep(1 * time.Second)
	a2.Play()
	time.Sleep(1 * time.Second)
	a3.Play()
	time.Sleep(1 * time.Second)
	a4.Play()
	time.Sleep(1 * time.Second)
}

func TestPan(t *testing.T) {
	a, err := wav.NewController().Wave(Sin(A4, 1, 2000))
	a2, err2 := a.MustCopy().Filter(filter.RightPan())
	a3, err3 := a.MustCopy().Filter(filter.LeftPan())
	assert.Nil(t, err)
	assert.Nil(t, err2)
	assert.Nil(t, err3)
	a.Play()
	time.Sleep(1 * time.Second)
	a2.Play()
	time.Sleep(1 * time.Second)
	a3.Play()
	time.Sleep(1 * time.Second)
	a5, _ := wav.NewController().Wave(Sin(A4, .1, 2000))
	for p := -1.0; p < 1; p += 0.04 {
		a6, _ := a5.MustCopy().Filter(filter.Pan(p))
		a6.Play()
		time.Sleep(100 * time.Millisecond)
	}
}

func testWave(t *testing.T, wave []byte) {
	a, err := wav.NewController().Wave(wave)
	assert.Nil(t, err)
	err = <-a.Play()
	assert.Nil(t, err)
	time.Sleep(2 * time.Second)
}
