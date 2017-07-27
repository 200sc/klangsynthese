package klangsynthese

import (
	"os"
	"strings"

	"github.com/200sc/klangsynthese/audio"
	"github.com/200sc/klangsynthese/flac"
	"github.com/200sc/klangsynthese/mp3"
	"github.com/200sc/klangsynthese/wav"
	"github.com/pkg/errors"
)

// LoadFile will parse the file name given and redirect to the appropriate
// subpackage depending on the ending.
func LoadFile(filename string) (audio.Audio, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to load file")
	}
	switch strings.ToLower(filename[len(filename)-4:]) {
	case ".wav":
		return wav.Load(f)
	case "flac":
		return flac.Load(f)
	case ".mp3":
		return mp3.Load(f)
	default:
		return nil, errors.New("Unsupported file type")
	}
}
