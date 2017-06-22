package synth

import (
	"testing"
	"time"

	"github.com/200sc/klangsynthese/wav"
	"github.com/stretchr/testify/assert"
)

func TestSinWav(t *testing.T) {
	testWave(t, Sin(A4, 2, 32))
}

func TestSquareWav(t *testing.T) {
	testWave(t, Square(A4, 2, 32))
}

func TestSawWav(t *testing.T) {
	testWave(t, Saw(A4, 2, 32))
}

func TestReverseSawWav(t *testing.T) {
	testWave(t, Reverse(Saw(A4, 2, 32)))
}

func TestTriangle(t *testing.T) {
	testWave(t, Triangle(A4, 2, 32))
}

func testWave(t *testing.T, wave []byte) {
	a, err := wav.NewController().Wave(wave)
	assert.Nil(t, err)
	err = <-a.Play()
	assert.Nil(t, err)
	time.Sleep(2 * time.Second)
}
