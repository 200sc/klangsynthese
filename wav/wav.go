package wav

import (
	"io"

	"encoding/binary"

	"github.com/200sc/klangsynthese/audio"
)

// def wav format
var format = audio.Format{
	SampleRate: 44100,
	Bits:       16,
	Channels:   2,
}

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (mc *Controller) Wave(b []byte) (audio.Audio, error) {
	return audio.EncodeBytes(
		audio.Encoding{
			Data:   b,
			Format: mc.Format(),
		})
}

func (mc *Controller) Load(r io.ReadCloser) (audio.Audio, error) {
	wav, err := ReadWavData(r)
	if err != nil {
		return nil, err
	}
	format := mc.Format()
	format.SampleRate = wav.SampleRate
	format.Channels = wav.NumChannels
	// This might also be WavData.ByteRate / 8, or BitsPerSample
	format.Bits = wav.BlockAlign * 8 / wav.NumChannels
	return audio.EncodeBytes(
		audio.Encoding{
			Data:   wav.Data,
			Format: format,
		})
}

func (mc *Controller) Save(r io.ReadWriter, a audio.Audio) error {
	return nil
}

func (mc *Controller) Format() audio.Format {
	return format
}

type WavData struct {
	bChunkID  [4]byte // B
	ChunkSize uint32  // L
	bFormat   [4]byte // B

	bSubchunk1ID  [4]byte // B
	Subchunk1Size uint32  // L
	AudioFormat   uint16  // L
	NumChannels   uint16  // L
	SampleRate    uint32  // L
	ByteRate      uint32  // L
	BlockAlign    uint16  // L
	BitsPerSample uint16  // L

	bSubchunk2ID  [4]byte // B
	Subchunk2Size uint32  // L
	Data          []byte  // L
}

func ReadWavData(r io.Reader) (WavData, error) {
	wav := WavData{}

	binary.Read(r, binary.BigEndian, &wav.bChunkID)
	binary.Read(r, binary.LittleEndian, &wav.ChunkSize)
	binary.Read(r, binary.BigEndian, &wav.bFormat)

	binary.Read(r, binary.BigEndian, &wav.bSubchunk1ID)
	binary.Read(r, binary.LittleEndian, &wav.Subchunk1Size)
	binary.Read(r, binary.LittleEndian, &wav.AudioFormat)
	binary.Read(r, binary.LittleEndian, &wav.NumChannels)
	binary.Read(r, binary.LittleEndian, &wav.SampleRate)
	binary.Read(r, binary.LittleEndian, &wav.ByteRate)
	binary.Read(r, binary.LittleEndian, &wav.BlockAlign)
	binary.Read(r, binary.LittleEndian, &wav.BitsPerSample)

	binary.Read(r, binary.BigEndian, &wav.bSubchunk2ID)
	binary.Read(r, binary.LittleEndian, &wav.Subchunk2Size)

	wav.Data = make([]byte, wav.Subchunk2Size)
	binary.Read(r, binary.LittleEndian, &wav.Data)

	return wav, nil
}
