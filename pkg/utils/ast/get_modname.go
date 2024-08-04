package gengast

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/mod/modfile"
)

// GetModuleNameFromGoModFile get's module name from go mod file in the directory provided
func GetModuleNameFromGoModFile(dir string) (string, error) {
	b, err := os.ReadFile(filepath.Join(dir, "go.mod"))
	if err != nil {
		return "", fmt.Errorf("error reading go.mod file: %w", err)
	}

	// parse the go mod file
	modfile, err := modfile.Parse("go.mod", b, nil)
	if err != nil {
		return "", fmt.Errorf("error parsing mod file: %w", err)
	}

	return modfile.Module.Mod.String(), nil
}
