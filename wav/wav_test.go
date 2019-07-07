package wav

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestBasicWav(t *testing.T) {
	fmt.Println("Running Basic Wav")
	f, err := os.Open("test.wav")
	fmt.Println(f)
	require.Nil(t, err)
	a, err := Load(f)
	require.Nil(t, err)
	err = <-a.Play()
	require.Nil(t, err)
	time.Sleep(4 * time.Second)
	// In addition to the error tests here, this should play noise
}
