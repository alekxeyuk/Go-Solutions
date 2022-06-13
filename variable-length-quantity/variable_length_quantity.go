package variablelengthquantity

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	maxbytes         = 8
	notthelast uint8 = 128
)

func EncodeVarint(input []uint32) []byte {
	out := []byte{}
	for _, value := range input {
		numbytes := []byte{}
		tmp := uint8(0)
		power := uint8(1)
		morethanone := false
		for i := 1; ; value >>= 1 {
			if i == maxbytes || value == 0 {
				if morethanone {
					tmp ^= notthelast
				} else {
					morethanone = !morethanone
				}
				numbytes = append([]byte{tmp}, numbytes...)
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
		out = append(out, numbytes...)
	}
	return out
}

func DecodeVarint(input []byte) ([]uint32, error) {
	var sb strings.Builder
	out := []uint32{}
	complete := false
	for _, value := range input {
		tmp := fmt.Sprintf("%08b", value)
		sb.WriteString(tmp[1:])
		if tmp[0] == '0' {
			p, _ := strconv.ParseUint(sb.String(), 2, 32)
			out = append(out, uint32(p))
			sb.Reset()
			complete = true
		} else {
			complete = false
		}
	}
	if !complete {
		return nil, errors.New("incomplete  sequence")
	}
	return out, nil
}
