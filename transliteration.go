package quranize

import (
	"strings"

	"github.com/alpancs/quranize/corpus"
)

// Transliteration helps Quranize to encode arabic into alphabet.
type Transliteration struct {
	hijaiyas       map[string][]string
	alphabetMaxLen int
}

var (
	base = []string{""}
)

// NewDefaultTransliteration returns new Transliteration using default mapping.
//
// Mapping: https://github.com/alpancs/quranize/blob/master/corpus/arabic_to_alphabet_clean.go#L3
func NewDefaultTransliteration() Transliteration {
	return NewTransliteration(corpus.ArabicToAlphabetClean)
}

// NewTransliteration returns new Transliteration.
func NewTransliteration(raw string) Transliteration {
	hijaiyas := make(map[string][]string)
	alphabetMaxLen := 0

	for _, line := range strings.Split(strings.TrimSpace(raw), "\n") {
		components := strings.Split(line, " ")
		arabic := components[0]
		for _, alphabet := range components[1:] {
			hijaiyas[alphabet] = append(hijaiyas[alphabet], arabic)

			length := len(alphabet)
			ending := alphabet[length-1]
			if ending == 'a' || ending == 'i' || ending == 'o' || ending == 'u' {
				alphabet = alphabet[:length-1] + alphabet[:length-1] + alphabet[length-1:]
			} else {
				alphabet += alphabet
			}
			hijaiyas[alphabet] = append(hijaiyas[alphabet], arabic)

			length = len(alphabet)
			if length > alphabetMaxLen {
				alphabetMaxLen = length
			}
		}
	}

	return Transliteration{hijaiyas, alphabetMaxLen}
}
