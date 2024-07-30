package utils

import (
	"fmt"
	"io"
	"os"
)

// IsDirEmpty checks if directory is empty or not
func IsDirEmpty(path string) (bool, error) {

	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return true, nil
	}

	if err != nil {
		return false, err
	}

  if !info.IsDir() {
    return false, fmt.Errorf("path is not a directory")
  }

	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	// Read directory entries.
	_, err = f.Readdir(1)
	if err == nil {
		return false, nil
	}

	if err == io.EOF {
		return true, nil
	}

	return false, err
}
