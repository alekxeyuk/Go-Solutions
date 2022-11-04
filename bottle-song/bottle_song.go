package bottlesong

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	VERSE_1   = " green bottle"
	VERSE_1_2 = " hanging on the wall,"
	VERSE_2   = "And if one green bottle should accidentally fall,"
	VERSE_3   = "There'll be "
	VERSE_4   = " green bottle"
	VERSE_4_2 = " hanging on the wall."
	NUMS      = []string{"no", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
)

func title(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToTitle(r)) + s[size:]
}

func verse(bottles_left int, sb *strings.Builder, song []string, songIndex int) int {
	if songIndex > 0 {
		song[songIndex] = ""
		songIndex++
	}
	sb.WriteString(title(NUMS[bottles_left]))
	sb.WriteString(VERSE_1)
	if bottles_left > 1 {
		sb.WriteString("s")
	}
	sb.WriteString(VERSE_1_2)
	song[songIndex] = sb.String()
	song[songIndex+1] = sb.String()
	songIndex += 2
	sb.Reset()

	sb.WriteString(VERSE_2)
	song[songIndex] = sb.String()
	songIndex++
	sb.Reset()

	sb.WriteString(VERSE_3)
	sb.WriteString(NUMS[bottles_left-1])
	sb.WriteString(VERSE_4)
	if bottles_left != 2 {
		sb.WriteString("s")
	}
	sb.WriteString(VERSE_4_2)
	song[songIndex] = sb.String()
	songIndex++
	sb.Reset()

	return songIndex
}

func Recite(startBottles, takeDown int) []string {
	song := make([]string, 5*takeDown-1)
	songIndex := 0
	var sb strings.Builder

	for i := startBottles; i > startBottles-takeDown; i-- {
		songIndex = verse(i, &sb, song, songIndex)
	}

	return song
}
