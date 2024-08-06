package utils

import (
	"strings"
)

func GetFileWithoutExt(file []string) []string {
	op := []string{}
	for _, k := range file {
		s := strings.Split(k, ".")
		op = append(op, s[0])
	}

	return op
}
