package wav

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/200sc/klangsynthese/synth"
	"github.com/stretchr/testify/assert"
)

func TestBasicWav(t *testing.T) {
	fmt.Println("Running Basic Wav")
	f, err := os.Open("test.wav")
	fmt.Println(f)
	assert.Nil(t, err)
	a, err := NewController().Load(f)
	assert.Nil(t, err)
	err = <-a.Play()
	assert.Nil(t, err)
	time.Sleep(4 * time.Second)
	// In addition to the error tests here, this should play noise
}

func TestSinWav(t *testing.T) {
	a, err := NewController().Wave(synth.Sin(150, 2, 64, 44100))
	assert.Nil(t, err)
	err = <-a.Play()
	assert.Nil(t, err)
	time.Sleep(2 * time.Second)
}
