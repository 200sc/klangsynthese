package filter

func setInt16(d []byte, i int, in int64) {
	for j := 0; j < 2; j++ {
		d[i+j] = byte(in & 255)
		in >>= 8
	}
}

func getInt16(d []byte, i int) (out int16) {
	var shift uint16
	for j := 0; j < 2; j++ {
		out += int16(d[i+j]) << shift
		shift += 8
	}
	return
}

func round(f float64) int64 {
	if f < 0 {
		return int64(f - .5)
	}
	return int64(f + .5)
}
