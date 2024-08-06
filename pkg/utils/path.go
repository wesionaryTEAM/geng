package utils

import (
	"fmt"
	"os"
	"path/filepath"
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

// Find the file path
func FindPathFromFile(startPath, file string) (string, error) {

	if startPath == "/" {
		return "", fmt.Errorf("reached root directory, %s not found", file)
	}
	absolutePath, err := filepath.Abs(startPath)
	if err != nil {
		return "", err
	}
	// Check if file exists in the current path
	if _, err := os.Stat(filepath.Join(absolutePath, file)); err == nil {
		return startPath, nil
	} else if !os.IsNotExist(err) {
		return "", err
	}

	parentDir := filepath.Dir(absolutePath)
	return FindPathFromFile(parentDir, file)

}
