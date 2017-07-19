package sf2

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/200sc/klangsynthese/font/riff"
	"github.com/stretchr/testify/assert"
)

func TestReadSf2(t *testing.T) {
	fl, err := os.Open("nolicenseforthis.sf2")
	assert.Nil(t, err)
	data, err := ioutil.ReadAll(fl)
	assert.Nil(t, err)
	// Todo: There should be a way to not have to readAll this
	r := riff.NewReader(data)
	r.Print()
}

type Data struct {
	riff.INFO
	Sdta
	Pdta
}

type VersionTag struct {
	Major, Minor uint16
}

type Sdta struct {
	Smpls [][]int16
	Sm24s [][]int8
}

type Pdta struct {
	Phdr []PresetHeader
	Pbag []PresetBag
	Pmod []ModList
	Pgen []GenList
	Inst []Inst
	Ibag []InstBag
	Imod []ModList
	Igen []GenList
	Shdr []Sample
}

type PresetHeader struct {
	AchPresetName              [20]int8
	Preset, Bank, PresetBagNdx uint16
	Library, Genre, Morphology uint32
}

type PresetBag struct {
	GenNdx, ModNdx uint16
}

type ModList struct {
	ModSrcOper    Modulator
	ModDestOper   Generator
	ModAmount     int16
	ModAmtSrcOper Modulator
	ModTransOper  Transform
}

type GenList struct {
	GenOper   Generator
	GenAmount uint16
}

type Inst struct {
	AchInstName [20]int8
	InstBagNdx  uint16
}

type InstBag struct {
	InstGenNdx, InstModNdx uint16
}

type Sample struct {
	AchSampleName                              [20]byte
	Start, End, StartLoop, EndLoop, SampleRate uint32
	ByOriginalKey                              byte
	ChCorrection                               int8
	SampleLink                                 uint16
	SampleType                                 SampleLink
}

type Modulator uint16

// Modulator enum
const ()

type Generator uint16

// Generator enum
const ()

type Transform uint16

// Transform enum
const ()

type SampleLink uint16

// SampleLink enum
const (
	MonoSample      SampleLink = 1
	RightSample     SampleLink = 2
	LeftSample      SampleLink = 4
	LinkedSample    SampleLink = 8
	RomMonoSample   SampleLink = 0x8001
	RomRightSample  SampleLink = 0x8002
	RomLeftSample   SampleLink = 0x8004
	RomLinkedSample SampleLink = 0x8008
)
