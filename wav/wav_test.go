package wav

import (
	"fmt"
	"os"
	"testing"
	"time"

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
}
