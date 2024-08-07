package models

import (
	"path/filepath"
	"strings"

	"github.com/mukezhz/geng/pkg"
	"github.com/mukezhz/geng/pkg/utils"
	"github.com/mukezhz/geng/templates"
)

type Service struct {
	Directory         string `mapstructure:"dir"`
	ProjectModuleName string `mapstructure:"mod"`

	ServiceType []string `mapstructure:"servicetype"`
}

// Fill fills up some empty data from project name if possible to do so
func (p *Service) AutoFill() {
	logger := pkg.GetLogger()

	if p.Directory == "" {
		p.Directory = "./" + p.ProjectModuleName
	}

	goModPath := getGoModPath(p.Directory, logger)
	p.Directory = goModPath

	if p.ProjectModuleName != "" {
		return
	}

	// if project module name is not provided, then fill it up from go.mod file
	goMod, err := GetModuleNameFromGoModFile(p.Directory)
	if err != nil {
		logger.Fatal("couldn't get module name from go.mod", "err", err)
	}
	p.ProjectModuleName = goMod.Module
}

func (p *Service) Validate() error {

	if len(p.ServiceType) == 0 {
		return nil
	}
	logger := pkg.GetLogger()

	// check directory structures are properly present or not
	pathToInfra := utils.IgnoreWindowsPath(filepath.Join(".", "templates", "wesionary", "service"))
	files, err := utils.ListEmbDir(templates.FS, pathToInfra)
	if err != nil {
		logger.Fatal("couldn't list infrastructure directories", "err", err)
	}
	filesWithoutExt := utils.GetFileWithoutExt(files)
	serviceNames := getServiceNameFromFiles(filesWithoutExt)

	// check infra choices provided matches from the available in the templates
	for _, file := range p.ServiceType {
		if exists := utils.StrInList(
			serviceNames,
			strings.ToLower(strings.TrimSpace(file))); !exists {
			logger.Fatalf("service type not found for type: %s", file)
		}
	}

	// check if mod files are present in the given directory
	if _, err := utils.FindPathFromFile(p.Directory, "go.mod"); err != nil {
		logger.Fatal("go.mod file not found", "err", err)
	}

	return nil
}

func getServiceNameFromFiles(files []string) []string {
	op := []string{}
	for _, k := range files {
		op = append(op, getNameOfService(k))
	}

	return op
}

func getNameOfService(s string) string {
	return strings.Split(s, "_")[0]
}
