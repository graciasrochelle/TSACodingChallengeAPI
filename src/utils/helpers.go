package utils

import (
	"bytes"
	"strings"
)

// NormalizePhoneNumber removes spaces and special chanarcters except plus(+)
func NormalizePhoneNumber(phone string) string {
	var buf bytes.Buffer
	for _, ch := range phone {
		if ch >= '0' && ch <= '9' || ch == '+' {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}

// NameToTitle capitalizes each word
func NameToTitleCase(name string) string {
	return strings.Title(strings.ToLower(name))
}
