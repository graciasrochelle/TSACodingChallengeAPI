package utils

import (
	"bytes"
	"regexp"
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

// NameToTitleCase capitalizes each word
func NameToTitleCase(name string) string {
	return strings.Title(strings.ToLower(name))
}

// IsPossibleNumber checkes if phone number is valid for AU
func IsPossibleNumber(number string) bool {
	match, _ := regexp.MatchString("(\\(+61\\)|\\+61|\\(0[1-9]\\)|0[1-9])?( ?-?[0-9]){6,9}", number)
	return match
}
