package gengast

import "unicode"

// IsModNameValid checks if the given module name is valid
func IsModNameValid(name string) bool {
	if name == "" {
		return false
	}

	for i, r := range name {
		if i == 0 && !unicode.IsLetter(r) && r != '_' {
			return false
		}
		if i > 0 && !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
			return false
		}
	}

	return true
}
