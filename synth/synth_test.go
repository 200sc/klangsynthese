package synth

import (
	"fmt"
	"testing"
	"time"

	"github.com/200sc/klangsynthese/audio/filter"
	"github.com/stretchr/testify/assert"
)

func TestSinWav(t *testing.T) {
	a, err := Int16.Sin()
	assert.Nil(t, err)
	a.Play()
	time.Sleep(Int16.PlayLength())
}

func TestSquareWav(t *testing.T) {
	a, err := Int16.Square()
	assert.Nil(t, err)
	a.Play()
	time.Sleep(Int16.PlayLength())
}

func TestSawWav(t *testing.T) {
	a, err := Int16.Saw()
	assert.Nil(t, err)
	a.Play()
	time.Sleep(Int16.PlayLength())
}

func TestTriangleWav(t *testing.T) {
	a, err := Int16.Triangle()
	assert.Nil(t, err)
	a.Play()
	time.Sleep(Int16.PlayLength())
}

func TestPulseWav(t *testing.T) {
	a, err := Int16.Pulse(8)()
	assert.Nil(t, err)
	a.Play()
	time.Sleep(Int16.PlayLength())
}

func TestVolume(t *testing.T) {
	a, _ := Int16.Sin()
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
	a, err := Int16.Sin()
	a2, err2 := a.MustCopy().Filter(filter.RightPan())
	a3, err3 := a.MustCopy().Filter(filter.LeftPan())
	assert.Nil(t, err)
	assert.Nil(t, err2)
	assert.Nil(t, err3)
	a.Play()
	fmt.Println(a.PlayLength())
	time.Sleep(a.PlayLength())
	a2.Play()
	time.Sleep(a2.PlayLength())
	a3.Play()
	time.Sleep(a3.PlayLength())
	a5, _ := Int16.Sin(Duration(100 * time.Millisecond))
	for p := -1.0; p < 1; p += 0.04 {
		a6, _ := a5.MustCopy().Filter(filter.Pan(p))
		a6.Play()
		time.Sleep(a6.PlayLength())
	}
}

func TestStop(t *testing.T) {
	a, _ := Int16.Sin()
	<-a.Play()
	_ = a.Stop()
	time.Sleep(1 * time.Second)
	// assert that sound was not heard or was only heard very briefly
}

func TestLoop(t *testing.T) {
	a, _ := Int16.Sin()
	a, _ = a.Filter(filter.LoopOn())
	<-a.Play()
	time.Sleep(3 * time.Second)
}

func TestModSampleRate(t *testing.T) {
	a, _ := Int16.Sin()
	a2, _ := a.MustCopy().Filter(filter.ModSampleRate(.5))
	a.Play()
	time.Sleep(1 * time.Second)
	a2.Play()
	time.Sleep(2 * time.Second)
}

func TestPitchShift(t *testing.T) {
	a, _ := Int16.Sin(Duration(1 * time.Second))
	a.Play()
	time.Sleep(a.PlayLength())
	a = a.MustFilter(filter.LowQualityShifter.PitchShift(0.5))
	a.Play()
	time.Sleep(a.PlayLength())
}

func TestSpeed(t *testing.T) {
	a, _ := Int16.Sin(Duration(1 * time.Second))
	a.Play()
	time.Sleep(a.PlayLength())
	a = a.MustFilter(filter.Speed(.5, filter.HighQualityShifter))
	a.Play()
	time.Sleep(a.PlayLength())
}
