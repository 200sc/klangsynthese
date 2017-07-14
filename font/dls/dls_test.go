package dls

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/200sc/klangsynthese/font/riffutil"
	"github.com/stretchr/testify/assert"
	"golang.org/x/image/riff"
)

func TestDLS(t *testing.T) {
	fl, err := os.Open("nolicenseforthis.dls")
	assert.Nil(t, err)
	typ, reader, err := riff.NewReader(fl)
	fmt.Println(riffutil.Header(typ))
	riffutil.DeepRead(reader)
}

func TestDLSUnmarshal(t *testing.T) {
	fl, err := os.Open("nolicenseforthis.dls")
	assert.Nil(t, err)
	dls := &DLS{}
	by, err := ioutil.ReadAll(fl)
	assert.Nil(t, err)
	err = riffutil.Unmarshal(by, dls)
	assert.Nil(t, err)
	//afmt := audio.Format{44100, 1, 16}
	fmt.Println(len(dls.Lins))
	fmt.Println("Version", dls.Vers)
	for _, ins := range dls.Lins {
		fmt.Println(ins.Insh)
		for _, rgn := range ins.Lrgn {
			fmt.Println(rgn)
		}
	}
	fmt.Println(dls.Info)
	//for _, w := range dls.Wvpl {
	//fmt.Println(w.Fmt)
	//a, err := afmt.Wave(w.Data)
	//assert.Nil(t, err)
	//a.Play()
	//time.Sleep(a.PlayLength())
	//}
}

type DLS struct {
	Dlid DLSID  `riff:"dlid"`
	Colh uint32 `riff:"colh"`
	Vers int64  `riff:"vers"`
	Lins []Ins  `riff:"lins"`
	Ptbl []byte `riff:"ptbl"` //PoolTable
	Wvpl []Wave `riff:"wvpl"`
	Info INFO   `riff:"INFO"`
}

type PoolTable struct {
	CbSize uint32
	CCues  uint32
	// CCues size
	PoolCues []PoolCue
}

type PoolCue struct {
	UlOffset uint32
}

type DLSID struct {
	UlData1 uint32
	UlData2 uint16
	UlData3 uint16
	AbData4 [8]byte
}

type Wave struct {
	Dlid DLSID     `riff:"dlid"`
	Guid []byte    `riff:"guid"`
	Wavu []byte    `riff:"wavu"`
	Fmt  PCMFormat `riff:"fmt "`
	Wavh []byte    `riff:"wavh"`
	Smpl []byte    `riff:"smpl"`
	Wsmp []byte    `riff:"wsmp"`
	Data []byte    `riff:"data"`
	Info INFO      `riff:"INFO"`
}

type PCMFormat struct {
	AudioFormat   uint16
	NumChannels   uint16
	SampleRate    uint32
	ByteRate      uint32
	BlockAlign    uint16
	BitsPerSample uint16
	WhoKnows      uint16 // There's buffer bytes?
}

type Ins struct {
	Dlid DLSID `riff:"dlid"`
	Insh Insh  `riff:"insh"`
	Lrgn []Rgn `riff:"lrgn"`
	Lart Art   `riff:"lart"`
	Info INFO  `riff:"INFO"`
}

type Insh struct {
	CRegions uint32
	Locale   MIDILOCALE
}

type MIDILOCALE struct {
	UlBank       uint32
	UlInstrument uint32
}

type Art struct {
	// Todo: art1 doesn't fit our unmarshaler's expectations, because
	// it's basically its own type of subchunk with two sizes following
	// 'art1' then a number of structs based on the second size
	Art1 []byte `riff:"art1"`
	// This []byte is really:
	// cbSize uint32
	// cConnectionBlocks uint32
	// ConnectionBlocks []ConnectionBlock
	// Also Art2, which is equivalent to Art1
}

type Rgn struct {
	Rgnh Rgnh     `riff:"rgnh"`
	Wsmp []byte   `riff:"wsmp"`
	Wlnk WaveLink `riff:"wlnk"`
	Lart Art      `riff:"lart"`
}

// Todo: figure out how to distinguish between rgnhs with and without ulLayer
type Rgnh struct {
	RangeKey      RGNRANGE
	RangeVelocity RGNRANGE
	FusOptions    uint16
	UsKeyGroup    uint16
	UsLayer       uint16 // This field is optional
}

type RGNRANGE struct {
	UsLow  uint16
	UsHigh uint16
}

type WaveLink struct {
	FusOptions   uint16
	UsPhaseGroup uint16
	UlChannel    uint32
	UlTableIndex uint32
}

type WaveSample struct {
	CbSize      uint32
	UsUnityNote uint16
	SFineTune   int16
	LGain       int32
	FulOptions  uint32
	// As for art, WaveSampleLoop is CSampleLoops long
	CSampleLoops   uint32
	WaveSampleLoop []WaveSampleLoop
}

type WaveSampleLoop struct {
	CbSize       uint32
	UlLoopType   uint32
	UlLoopStart  uint32
	UlLoopLength uint32
}

type INFO struct {
	IARL string `riff:"IARL"`
	IART string `riff:"IART"`
	ICMS string `riff:"ICMS"`
	ICMT string `riff:"ICMT"`
	ICOP string `riff:"ICOP"`
	ICRD string `riff:"ICRD"`
	IENG string `riff:"IENG"`
	IGNR string `riff:"IGNR"`
	IKEY string `riff:"IKEY"`
	IMED string `riff:"IMED"`
	INAM string `riff:"INAM"`
	IPRD string `riff:"IPRD"`
	ISBJ string `riff:"ISBJ"`
	ISFT string `riff:"ISFT"`
	ISRC string `riff:"ISRC"`
	ISRF string `riff:"ISRF"`
	ITCH string `riff:"ITCH"`
}
