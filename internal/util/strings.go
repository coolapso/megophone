package util

import (
	"strings"
)

func MaskString(s string) string {
	l := len(s)
	if l <= 3 {
		return s
	}

	return strings.Repeat("*", l-3) + s[l-3:]
}

func CleanString(s string) string {
	s = strings.ReplaceAll(s, "\\n", "\n")
	s = strings.ReplaceAll(s, "\\t", "\t")

	return s
}


func IsXLenght(s string) bool {
	return len(CleanString(s)) <= 280
}

func IsToothLenght(s string) bool {
	return len(CleanString(s)) <= 500
}
