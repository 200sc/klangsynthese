package dls

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/200sc/klangsynthese/audio"
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

func TestDLSUnmarshall(t *testing.T) {
	fl, err := os.Open("nolicenseforthis.dls")
	assert.Nil(t, err)
	dls := &DLS{}
	by, err := ioutil.ReadAll(fl)
	assert.Nil(t, err)
	err = riffutil.Unmarshal(by, dls)
	assert.Nil(t, err)
	afmt := audio.Format{44100, 1, 16}
	for i := range dls.Wvpl {
		a, err := afmt.Wave(dls.Wvpl[i].Data)
		assert.Nil(t, err)
		a.Play()
		time.Sleep(a.PlayLength())
	}
}

type DLS struct {
	Dlid []byte `riff:"dlid"`
	Colh []byte `riff:"colh"`
	Vers []byte `riff:"vers"`
	Lins []Ins  `riff:"lins"`
	Ptbl []byte `riff:"ptbl"`
	Wvpl []Wave `riff:"wvpl"`
	Info INFO   `riff:"INFO"`
}

type Wave struct {
	Guid []byte `riff:"guid"`
	Wavu []byte `riff:"wavu"`
	Fmt  []byte `riff:"fmt "`
	Wavh []byte `riff:"wavh"`
	Smpl []byte `riff:"smpl"`
	Wsmp []byte `riff:"wsmp"`
	Data []byte `riff:"data"`
	Info INFO   `riff:"INFO"`
}

type Ins struct {
	Insh []byte `riff:"insh"`
	Lrgn []Rgn  `riff:"lrgn"`
	Lart Art    `riff:"lart"`
	Info INFO   `riff:"INFO"`
}

type Art struct {
	Art1 []byte `riff:"art1"`
}

type Rgn struct {
	Rgnh []byte `riff:"rgnh"`
	Wsmp []byte `riff:"wsmp"`
	Wlnk []byte `riff:"wlnk"`
}

type INFO struct {
	ICMT []byte `riff:"ICMT"`
	ICOP []byte `riff:"ICOP"`
	IENG []byte `riff:"IENG"`
	INAM []byte `riff:"INAM"`
	ISBJ []byte `riff:"ISBJ"`
}
