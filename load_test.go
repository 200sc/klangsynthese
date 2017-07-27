package klangsynthese

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// These tests do not test that we can play the given audio, see the
// respective directory for each file type for those tests. These
// tests just confirm that we can load the files.

func TestLoadMp3(t *testing.T) {
	_, err := LoadFile("mp3/test.mp3")
	assert.Nil(t, err)
}

func TestLoadWav(t *testing.T) {
	_, err := LoadFile("wav/test.wav")
	assert.Nil(t, err)
}

func TestLoadFlac(t *testing.T) {
	_, err := LoadFile("flac/test.flac")
	assert.Nil(t, err)
}
