package riffutil

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"reflect"

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
			fmt.Print(prefix, Header(typ), "Length:", l)
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
					fmt.Println(" Content:", string(b))
				}
			} else {
				fmt.Println(" Long Content")
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

func Unmarshal(data []byte, v interface{}) error {
	var d decodeState
	d.reader = bytes.NewReader(data)
	return d.unmarshal(v)
}

type decodeState struct {
	reader *bytes.Reader
}

func (ds *decodeState) unmarshal(v interface{}) error {
	// Mirrors json.unmarshal
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("Invalid Unmarshal Struct")
	}
	// The first ID in the riff should be RIFF
	id, err := ds.nextID()
	if err != nil {
		return err
	}
	if id != "RIFF" {
		return errors.New("RIFF format must begin with RIFF")
	}
	ln, err := ds.nextLen()
	if err != nil {
		return err
	}
	// The next ID identifies this file type. We don't want it.
	_, err = ds.nextID()
	if err != nil {
		return err
	}
	_, err = ds.chunks(reflect.Indirect(rv), int(ln))
	return err
}

func (ds *decodeState) nextID() (string, error) {
	id := make([]byte, 4)
	l, err := ds.reader.Read(id)
	if l != 4 || err != nil {
		return "", errors.New("RIFF missing expected ID")
	}
	return string(id), nil
}

func (ds *decodeState) nextIdLen() (string, uint32, bool, error) {
	id, err := ds.nextID()
	if err != nil {
		return "", 0, false, err
	}
	ln, err := ds.nextLen()
	if err != nil {
		return "", 0, false, err
	}
	return id, ln, id == "LIST", nil
}

func (ds *decodeState) nextLen() (uint32, error) {
	var ln uint32
	err := binary.Read(ds.reader, binary.LittleEndian, &ln)
	if err != nil {
		return ln, errors.New("RIFF missing expected length")
	}
	return ln, nil
}

func (ds *decodeState) chunks(rv reflect.Value, inLength int) (reflect.Value, error) {
	// Find chunkId in rv
	// If it can't be found, ignore it as a value the user does not want
	switch rv.Kind() {
	case reflect.Struct:
		return rv, ds.structChunks(rv, inLength)
	case reflect.Slice:
		return ds.sliceChunks(rv, inLength)
	default:
		return reflect.Value{}, errors.New("Unsupported unmarshal type")
	}
}

func (ds *decodeState) sliceChunks(rv reflect.Value, inLength int) (reflect.Value, error) {

	slTy := rv.Type()
	ty := slTy.Elem()
	newSlice := reflect.MakeSlice(slTy, 0, 10000)
	for inLength > 0 {
		_, ln, isList, err := ds.nextIdLen()
		if err != nil {
			return reflect.Value{}, err
		}
		if !isList {
			return reflect.Value{}, errors.New("Slice structs need to be LISTs")
		}
		ln -= 4
		inLength -= 4
		if inLength <= 0 {
			break
		}
		_, err = ds.nextID()
		if err != nil {
			return reflect.Value{}, err
		}

		inLength -= 8
		if inLength <= 0 {
			break
		}
		newStruct := reflect.New(ty)
		err = ds.structChunks(reflect.Indirect(newStruct), int(ln))
		newSlice = reflect.Append(newSlice, reflect.Indirect(newStruct))
		if ln%2 != 0 {
			ds.reader.ReadByte()
			inLength--
		}
		inLength -= int(ln)
	}
	return newSlice, nil
}

// structChunks reads chunks and matches them to fields on rv (which is a struct)
// structChunks sets the fields of rv to be the ouput it gets
func (ds *decodeState) structChunks(rv reflect.Value, inLength int) error {
	chunkId, ln, isList, err := ds.nextIdLen()
	if err != nil {
		return err
	}
	if isList {
		ln -= 4
		inLength -= 4
		chunkId, err = ds.nextID()
		if err != nil {
			return err
		}
	}
	inLength -= 8
	ty := reflect.TypeOf(rv.Interface())
	fields := make([]reflect.Value, rv.NumField())
	fieldTags := make([]reflect.StructTag, rv.NumField())
	for i := range fields {
		fields[i] = rv.Field(i)
		fieldTags[i] = ty.Field(i).Tag
	}
	i := 0
	for inLength > 0 {
		tag := fieldTags[i].Get("riff")
		//spew.Dump(fields[i])
		if tag == chunkId {
			// get contents from recursive call
			var content reflect.Value
			if isList {
				content, err = ds.chunks(fields[i], int(ln))
			} else {
				content, err = ds.fieldValue(fields[i], ln)
			}
			if err != nil {
				return err
			}
			inLength -= int(ln)

			fields[i].Set(content)
			// if length is odd read one more
			if ln%2 != 0 {
				ds.reader.ReadByte()
				inLength--
			}
			if inLength <= 0 {
				return nil
			}
			// next id
			chunkId, ln, isList, err = ds.nextIdLen()
			if err != nil {
				return err
			}
			if isList {
				ln -= 4
				inLength -= 4
				chunkId, err = ds.nextID()
				if err != nil {
					return err
				}
			}
			inLength -= 8
			i = -1
		}
		if inLength <= 0 {
			return nil
		}
		i++
		if i >= len(fields) {
			// Skip this id
			// if length is odd read one more
			if ln%2 != 0 {
				ln++
			}
			_, err = ds.reader.Seek(int64(ln), io.SeekCurrent)
			if err != nil {
				return err
			}
			inLength -= int(ln)
			if inLength <= 0 {
				return nil
			}
			// next id
			chunkId, ln, isList, err = ds.nextIdLen()
			if err != nil {
				return err
			}
			if isList {
				ln -= 4
				inLength -= 4
				chunkId, err = ds.nextID()
				if err != nil {
					return err
				}
			}
			inLength -= 8
			i = 0
		}
	}
	return nil
}

func (ds *decodeState) fieldValue(rv reflect.Value, ln uint32) (reflect.Value, error) {
	switch rv.Kind() {
	case reflect.Slice:
		switch rv.Type().Elem().Kind() {
		case reflect.Uint8:
			data := make([]byte, ln)
			n, err := ds.reader.Read(data)
			if n != int(ln) {
				return reflect.Value{}, errors.New("Insufficient data found in RIFF data block")
			}
			return reflect.ValueOf(data), err
		default:
			return reflect.Value{}, errors.New("Unsupported type in input struct")
		}
	}
	return reflect.Value{}, nil
}

// Struct -> Fields
//
// Fields -> Field
//
// Field -> Struct
//       -> Slice
//
// Slice -> []byte
//       -> Struct
