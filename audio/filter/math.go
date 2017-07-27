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

func getFloat64(d []byte, i int, byteDepth uint16) float64 {
	switch byteDepth {
	case 1:
		return float64(int8(d[i])) / 128.0
	case 2:
		return float64(getInt16(d, i)) / 32768.0
	}
	return 0.0
}

func setInt16_f64(d []byte, i int, in float64) {
	setInt16(d, i, int64(in*32768))
}

func round(f float64) int64 {
	if f < 0 {
		return int64(f - .5)
	}
	return int64(f + .5)
}

// wrapMod will wrap the input i until it fits between [0, length).
// works for positive and negative numbers.
func wrapMod(i int, length int) int {
	if i > 0 {
		return i % length
	}
	for i < 0 {
		i += length
	}
	return i
}
