package say

import (
	"strings"
)

var (
	ones      = [10]string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	tens      = [10]string{"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	teens     = [10]string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	thousands = [4]string{"", " thousand", " million", " billion"}
)

func divMod(n int64, d int64) (div int64, mod int64) {
	div, mod = n/d, n%d
	return
}

func toText(n int64) string {
	if n < 10 {
		return ones[n]
	} else if n < 20 {
		return teens[n-10]
	} else {
		if n%10 > 0 {
			return tens[n/10] + "-" + ones[n%10]
		} else {
			return tens[n/10]
		}
	}
}

func Say(n int64) (string, bool) {
	if n < 0 || n > 999_999_999_999 {
		return "", false
	}
	if n == 0 {
		return "zero", true
	}
	out := ""
	m, thousandsIndex := int64(0), uint8(0)
	for n > 0 {
		n, m = divMod(n, 1000)
		newPart := ""
		if m > 99 {
			d, m2 := divMod(m, 100)
			newPart = toText(d) + " hundred " + toText(m2)
		} else if m > 0 {
			newPart = toText(m)
		}
		if thousandsIndex > 0 && m > 0 {
			newPart = newPart + thousands[thousandsIndex]
		}
		out = newPart + " " + out
		thousandsIndex++
	}
	return strings.TrimSpace(out), true
}
