package utils

import (
	. "strings"
)

func IsEmpty(str string) bool {
	if str = TrimSpace(ToLower(str)); str == "" {
		return true
	}
	return false
}
