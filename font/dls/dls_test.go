package dls

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/200sc/klangsynthese/audio"
	"github.com/200sc/klangsynthese/audio/filter"
	"github.com/200sc/klangsynthese/font/riff"
	"github.com/stretchr/testify/assert"
)

// SanbikiSCC.dls is the original soundfont from LA-MULANA, and is
// provided with Nigoro's consent. SanbikiSCC was written by
// Samieru: https://twitter.com/samieru_nigoro

func TestDLSPrint(t *testing.T) {
	fl, err := os.Open("SanbikiSCC.dls")
	assert.Nil(t, err)
	data, err := ioutil.ReadAll(fl)
	assert.Nil(t, err)
	// Todo: There should be a way to not have to readAll this
	r := riff.NewReader(data)
	r.Print()
	//fmt.Println(s)
}

func TestDLSUnmarshal(t *testing.T) {
	fl, err := os.Open("SanbikiSCC.dls")
	assert.Nil(t, err)
	dls := &DLS{}
	by, err := ioutil.ReadAll(fl)
	assert.Nil(t, err)
	err = riff.Unmarshal(by, dls)
	assert.Nil(t, err)
	afmt := audio.Format{44100, 1, 16}
	fmt.Println(len(dls.Lins))
	fmt.Println("Version", dls.Vers)
	for _, ins := range dls.Lins {
		fmt.Println(ins.Insh)
		for _, rgn := range ins.Lrgn {
			fmt.Println(rgn)
		}
	}
	fmt.Println(dls.INFO)
	// for i, w := range dls.Wvpl {
	// 	fmt.Println(i, w.Fmt)
	// 	a, err := afmt.Wave(w.Data)
	// 	assert.Nil(t, err)
	// 	a.Play()
	// 	time.Sleep(a.PlayLength())
	// }
	wv := dls.Wvpl[25]
	shft, err := filter.NewFFTShifter(32, 16)
	assert.Nil(t, err)
	fmt.Println(len(wv.Data))
	a, _ := afmt.Wave(wv.Data)
	a.Play()
	time.Sleep(a.PlayLength() * 12)
	for i := .5; i <= 2.0; i += .5 {
		a2 := a.MustCopy().MustFilter(shft.PitchShift(i))
		a2.Play()
		time.Sleep(a2.PlayLength() * 12)
	}
}
