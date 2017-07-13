package dls

import (
	"fmt"
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

type DLS struct {
	Dlid []byte `riff:dlid`
	Colh []byte `riff:colh`
	Vers []byte `riff:vers`
	Lins []Ins  `riff:lins`
	Info INFO   `riff:INFO`
}

type Ins struct {
	Insh []byte   `riff:insh`
	Lrgn []Rgn    `riff:lrgn`
	Lart [][]byte `riff:lart`
	Info INFO     `riff:INFO`
}

type Rgn struct {
	Rgnh [14]byte `riff:rgnh` // ???? these numbers might not be constant
	Wsmp []byte   `riff:wsmp`
	Wlnk [12]byte `riff:wlnk`
}

type INFO struct {
	ICMT string `riff:ICMT`
	ICOP string `riff:ICOP`
	IENG string `riff:IENG`
	INAM string `riff:INAM`
	ISBJ string `riff:ISBJ`
}
