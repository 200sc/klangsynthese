package manip

func BytesToF64(data []byte, channels, bitRate uint16, channel int) []float64 {
	byteDepth := bitRate / 8
	out := make([]float64, len(data)/int(channels*byteDepth))
	for i := channel * int(byteDepth); i < len(data); i += int(byteDepth * 2) {
		out[i/int(byteDepth*2)] = GetFloat64(data, i, byteDepth)
	}
	return out
}

