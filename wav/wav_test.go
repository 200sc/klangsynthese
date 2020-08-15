package wav

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestBasicWav(t *testing.T) {
	fmt.Println("Running Basic Wav")
	f, err := os.Open("test.wav")
	fmt.Println(f)
	if err != nil {
		t.Fatal("expected open err to be nil, was", err)
	}
	a, err := Load(f)
	if err != nil {
		t.Fatal("expected load err to be nil, was", err)
	}
	err = <-a.Play()
	if err != nil {
		t.Fatal("expected play err to be nil, was", err)
	}
	time.Sleep(4 * time.Second)
	// In addition to the error tests here, this should play noise
}
