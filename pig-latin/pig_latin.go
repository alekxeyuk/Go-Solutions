package piglatin

import "strings"

var lst = map[string]bool{"rh": true, "th": true, "sq": true, "qu": true, "sh": true, "gl": true, "ch": true, "ph": true, "tr": true, "br": true, "fr": true, "bl": true, "gr": true, "st": true, "sl": true, "cl": true, "pl": true, "fl": true}
var strt = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
var tri = map[string]bool{"squ": true, "thr": true, "sch": true}

func isLetter(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func isWord(s string) bool {
	for _, c := range s {
		if !isLetter(c) {
			return false
		}
	}
	return true
}

func Sentence(sentence string) string {
	words := strings.Fields(sentence)
	for i, word := range words {
		if strt[word[0]] || word[:2] == "yt" || word[:2] == "xr" {
			words[i] = word + "ay"
		} else if len(word) >= 3 && tri[word[:3]] {
			words[i] = word[3:] + word[:3] + "ay"
		} else if lst[word[:2]] {
			words[i] = word[2:] + word[:2] + "ay"
		} else if isWord(word) {
			words[i] = word[1:] + word[:1] + "ay"
		}
	}
	return strings.Join(words, " ")
}
