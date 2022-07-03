package utils

import "strings"

func Normalize(str string) string {
	return strings.TrimSpace(strings.ToLower(str))
}
