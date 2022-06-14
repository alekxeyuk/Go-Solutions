package variablelengthquantity

import (
	"errors"
)

const (
	chunkSize uint8 = 7
	eightBit  uint8 = 128
)

func EncodeVarint(input []uint32) []byte {
	var out []byte
	for _, value := range input {
		if value == 0 {
			out = append(out, 0x0)
		} else {
			var numBytes []byte
			for i := 0; value > 0; value >>= chunkSize {
				tmp := byte(value & 127)
				if i++; i > 1 {
					tmp ^= eightBit
				}
				numBytes = append([]byte{tmp}, numBytes...)
			}
			out = append(out, numBytes...)
		}
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
			temp <<= chunkSize
			complete = false
		}
	}
	if !complete {
		return nil, errors.New("incomplete  sequence")
	}
	return out, nil
}
