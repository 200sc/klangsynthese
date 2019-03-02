package flac

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBasicFlac(t *testing.T) {
	fmt.Println("Opening Basic Flac")
	f, err := os.Open("test2.flac")
	fmt.Println(f)
	assert.Nil(t, err)
	a, err2 := Load(f)
	fmt.Println(a)
	assert.Nil(t, err2)
	fmt.Println("Now playing")
	err = <-a.Play()
	assert.Nil(t, err)
	time.Sleep(8 * time.Second)
	// In addition to the error tests here, this should play noise
}
