package models

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

type GoMod struct {
	Module    string
	GoVersion string
}

func GetModuleNameFromGoModFile(goModPath string) (GoMod, error) {
	file, err := os.Open(filepath.Join(goModPath, "go.mod"))
	goMod := GoMod{}
	if err != nil {
		return goMod, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			// Extract module name
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				goMod.Module = parts[1]
			}
		} else if strings.HasPrefix(line, "go ") {
			// Extract Go version
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				goMod.GoVersion = parts[1]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return goMod, err
	}

	return goMod, nil
}
