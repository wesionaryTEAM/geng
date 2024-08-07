package models

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mukezhz/geng/pkg"
	"github.com/mukezhz/geng/pkg/utils"
	gengast "github.com/mukezhz/geng/pkg/utils/ast"
)

type Module struct {
	// Directory directory to generate module in, project root directory
	Directory string `mapstructure:"dir"`

	// PackageName module name to generate
	PackageName string `mapstructure:"name"`

	// ProjectModuleName project module name
	ProjectModuleName string

	// ModuleName model name for the template
	ModuleName string
}

func (m *Module) AutoFill() {

	logger := pkg.GetLogger()

	if m.Directory == "" {
		m.Directory = "./" // current directory
	}
	goModPath := getGoModPath(m.Directory, logger)
	m.Directory = goModPath

	if m.ProjectModuleName != "" {
		return
	}

	modName, err := gengast.GetModuleNameFromGoModFile(m.Directory)
	if err != nil {
		logger.Fatal("couldn't get module name from go.mod", "err", err)
	}

	m.ProjectModuleName = modName

	if len(m.PackageName) > 1 {
		m.ModuleName = strings.ToUpper(m.PackageName[0:1]) + m.PackageName[1:]
	}
}

func (m *Module) Validate() error {

	if m.PackageName == "" {
		return fmt.Errorf("module name is required")
	}

	if len(m.PackageName) < 2 {
		return fmt.Errorf("package name length is invalid")
	}

	if !gengast.IsModNameValid(m.PackageName) {
		return fmt.Errorf("module name is invalid")
	}
	// check if domain directory exists [project is not initialized]
	if ok, err := utils.IsDirEmpty(filepath.Join(m.Directory, "domain")); ok && err == nil {
		return fmt.Errorf("project is not initialized")
	}

	// check if module already exists in the directory
	if ok, err := utils.IsDirEmpty(filepath.Join(m.Directory, "domain", m.PackageName)); !ok || err != nil {

		if err != nil {
			return fmt.Errorf("error checking if module exists, err: %w", err)
		}

		return fmt.Errorf("module already exists")
	}

	return nil
}
