package utils

import (
	"os"
	"runtime"
	"strings"
)

func IgnoreWindowsPath(p string) string {
	if runtime.GOOS == "windows" {
		return strings.ReplaceAll(p, "\\", "/")
	}
	return p
}

// WriteToPath writes content to path
func WriteToPath(path, content string) error {
	return os.WriteFile(path, []byte(content), os.ModePerm)
}
