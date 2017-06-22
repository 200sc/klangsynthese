package synth

type Pitch uint16

const (
	C0  Pitch = 16
	C0s       = 17
	D0b       = 17
	D0        = 18
	D0s       = 20
	E0b       = 20
	E0        = 21
	F0        = 22
	F0s       = 23
	G0b       = 23
	G0        = 25
	G0s       = 26
	A0b       = 26
	A0        = 28
	A0s       = 29
	B0b       = 29
	B0        = 31
	C1        = 33
	C1s       = 35
	D1b       = 35
	D1        = 37
	D1s       = 39
	E1b       = 39
	E1        = 41
	F1        = 44
	F1s       = 46
	G1b       = 46
	G1        = 49
	G1s       = 52
	A1b       = 52
	A1        = 55
	A1s       = 58
	B1b       = 58
	B1        = 62
)

func (p Pitch) UpOctave() Pitch {
	return p * 2 // If only
}

func (p Pitch) DownOctave() Pitch {
	return p / 2 // If only
}
