package diamond

import (
	"errors"
	"strings"
)

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", errors.New("wrong input character")
	}

	sliceSize := int(char - 'A')
	tempSlice := make([]string, sliceSize*2+1)
	var sb strings.Builder

	for i := 0; i <= sliceSize; i++ {
		sb.WriteString(strings.Repeat(" ", sliceSize-i))
		sb.WriteByte(byte(i) + 'A')
		if i != 0 && i != sliceSize+1 {
			sb.WriteString(strings.Repeat(" ", 2*i-1))
			sb.WriteByte(byte(i) + 'A')
		}
		sb.WriteString(strings.Repeat(" ", sliceSize-i))
		tempSlice[i] = sb.String()
		sb.Reset()
	}

	for i := 0; i <= sliceSize; i++ {
		tempSlice[sliceSize+i] = tempSlice[sliceSize-i]
	}

	return strings.Join(tempSlice, "\n"), nil
}
