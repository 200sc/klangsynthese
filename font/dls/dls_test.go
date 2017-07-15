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