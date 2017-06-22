package synth

import (
	"testing"
	"time"

	"github.com/200sc/klangsynthese/wav"
	"github.com/stretchr/testify/assert"
)

func TestSinWav(t *testing.T) {
	a, err := wav.NewController().Wave(Sin(A4, 2, 64))
	assert.Nil(t, err)
	err = <-a.Play()
	assert.Nil(t, err)
	time.Sleep(2 * time.Second)
}

func TestSquareWav(t *testing.T) {
	a, err := wav.NewController().Wave(Square(C4, 2, 128))
	assert.Nil(t, err)
	err = <-a.Play()
	assert.Nil(t, err)
	time.Sleep(2 * time.Second)
}

// Not yet!
func TestSawWav(t *testing.T) {
	a, err := wav.NewController().Wave(Saw(A4, 2, 32))
	assert.Nil(t, err)
	err = <-a.Play()
	assert.Nil(t, err)
	time.Sleep(2 * time.Second)
}
