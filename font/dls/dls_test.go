package dls

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/200sc/klangsynthese/font/riff"
	"github.com/stretchr/testify/assert"
)

func TestDLSPrint(t *testing.T) {
	fl, err := os.Open("nolicenseforthis.dls")
	assert.Nil(t, err)
	data, err := ioutil.ReadAll(fl)
	assert.Nil(t, err)
	// Todo: There should be a way to not have to readAll this
	r := riff.NewReader(data)
	r.Print()
	//fmt.Println(s)
}

func TestDLSUnmarshal(t *testing.T) {
	fl, err := os.Open("nolicenseforthis.dls")
	assert.Nil(t, err)
	dls := &DLS{}
	by, err := ioutil.ReadAll(fl)
	assert.Nil(t, err)
	err = riff.Unmarshal(by, dls)
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
	fmt.Println(dls.INFO)
	//for _, w := range dls.Wvpl {
	//fmt.Println(w.Fmt)
	//a, err := afmt.Wave(w.Data)
	//assert.Nil(t, err)
	//a.Play()
	//time.Sleep(a.PlayLength())
	//}
}
