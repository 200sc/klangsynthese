package sf2

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/image/riff"
)

type RiffHeader riff.FourCC

func (rh RiffHeader) String() string {
	return string(rh[0]) + string(rh[1]) + string(rh[2]) + string(rh[3])
}

func TestReadSf2(t *testing.T) {
	fl, err := os.Open("nolicenseforthis.sf2")
	assert.Nil(t, err)
	typ, reader, err := riff.NewReader(fl)
	fmt.Println(RiffHeader(typ))
	DeepReadSf2(reader)
}

func DeepReadSf2(r *riff.Reader) {
	deepReadSf2(r, " ")
}
func deepReadSf2(r *riff.Reader, prefix string) {
	var err error
	var typ riff.FourCC
	var l uint32
	var data io.Reader
	for err == nil {
		typ, l, data, err = r.Next()
		if err == nil {
			fmt.Println(prefix, RiffHeader(typ), "Length:", l)
			//fmt.Println(prefix, data)
			if typ == riff.LIST {
				typ2, r2, err2 := riff.NewListReader(l, data)
				if err2 == nil {
					fmt.Println(prefix+"  ", RiffHeader(typ2))
					deepReadSf2(r2, prefix+"    ")
				} else {
					fmt.Println(prefix, err2)
				}
			}
		}
	}
	if err != io.EOF {
		fmt.Println(prefix, err)
	}
}
