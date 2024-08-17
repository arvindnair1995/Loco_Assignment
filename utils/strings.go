package utils

import "strings"

func TransformString(s string) string {
	s = strings.TrimSpace(s)
	return strings.ToUpper(s)
}
