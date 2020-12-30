package app

import (
	"strings"
)

func mixedCaps(s string) string {
	s = strings.ToLower(s)
	words := strings.Split(s, " ")
	for i := 1; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}
