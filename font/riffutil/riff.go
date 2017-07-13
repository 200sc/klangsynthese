package riffutil

import (
	"fmt"
	"io"

	"golang.org/x/image/riff"
)

func DeepRead(r *riff.Reader) {
	deepRead(r, " ")
}
func deepRead(r *riff.Reader, prefix string) {
	var err error
	var typ riff.FourCC
	var l uint32
	var data io.Reader
	for err == nil {
		typ, l, data, err = r.Next()
		if err == nil {
			fmt.Println(prefix, Header(typ), "Length:", l)
			//fmt.Println(prefix, data)
			if typ == riff.LIST {
				typ2, r2, err2 := riff.NewListReader(l, data)
				if err2 == nil {
					fmt.Println(prefix+"  ", Header(typ2))
					deepRead(r2, prefix+"    ")
				} else {
					fmt.Println(prefix, err2)
				}
			} else if l < 40 {
				b := make([]byte, l)
				n, err := data.Read(b)
				if err != nil || n != int(l) {
					fmt.Println("Error: ", err)
				} else {
					fmt.Println(prefix, "Content:", string(b))
				}
			}
		}
	}
	if err != io.EOF {
		fmt.Println(prefix, err)
	}
}

type Header riff.FourCC

func (rh Header) String() string {
	return string(rh[0]) + string(rh[1]) + string(rh[2]) + string(rh[3])
}
