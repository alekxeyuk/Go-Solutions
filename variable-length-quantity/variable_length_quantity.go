package variablelengthquantity

import (
	"errors"
)

const (
	maxBits        = 8
	eightBit uint8 = 128
)

func EncodeVarint(input []uint32) []byte {
	var out []byte
	for _, value := range input {
		var numBytes []byte
		tmp := uint8(0)
		power := uint8(1)
		morethanone := false
		for i := 1; ; value >>= 1 {
			if i == maxBits || value == 0 {
				if morethanone {
					tmp ^= eightBit
				} else {
					morethanone = !morethanone
				}
				numBytes = append([]byte{tmp}, numBytes...)
				tmp = 0
				power = 1
				i = 1
			}
			if value == 0 {
				break
			} else if value&1 != 0 {
				tmp ^= power
			}
			power *= 2
			i++
		}
		out = append(out, numBytes...)
	}
	return out
}

func DecodeVarint(input []byte) ([]uint32, error) {
	var out []uint32
	temp := uint32(0)
	complete := false
	for _, value := range input {
		if value&eightBit == 0 {
			out = append(out, temp|uint32(value))
			complete = true
			temp = 0
		} else {
			temp |= uint32(value ^ eightBit)
			temp <<= 7
			complete = false
		}
	}
	if !complete {
		return nil, errors.New("incomplete  sequence")
	}
	return out, nil
}
