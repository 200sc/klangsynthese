package mp3

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBasicMp3(t *testing.T) {
	fmt.Println("Running Basic Mp3")
	f, err := os.Open("nolicenseforthis_test.mp3")
	fmt.Println(f)
	assert.Nil(t, err)
	a, err := NewController().Load(f)
	assert.Nil(t, err)
	err = <-a.Play()
	assert.Nil(t, err)
	fmt.Println("Starting playing")
	time.Sleep(10 * time.Second)
}
