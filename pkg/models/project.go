package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mukezhz/geng/pkg/utils"
)

type Project struct {
	PackageName string `mapstructure:"pkg"`

	ProjectDescription string `mapstructure:"desc"`
	ProjectModuleName  string `mapstructure:"mod"`
	ProjectName        string `mapstructure:"name"`
	Author             string `mapstructure:"author"`

	GoVersion string `mapstructure:"goversion"`
	Directory string `mapstructure:"dir"`
}

// Fill fills up some empty data from project name if possible to do so
func (p *Project) AutoFill() {
	if p.ProjectName == "" {
		return
	}

	projSplit := strings.Split(strings.ToLower(p.ProjectName), " ")
	pS := []string{}
	for _, k := range projSplit {
		s := strings.TrimSpace(k)
		if s != "" {
			pS = append(pS, s)
		}
	}

	projSplit = pS

	if p.ProjectModuleName == "" {
		p.ProjectModuleName = strings.Join(projSplit, "_")
	}

	if p.PackageName == "" {
		p.PackageName = strings.Join(projSplit, "_")
	}

	if p.Directory == "" {
		p.Directory = "./" + p.PackageName
	}

}

func (p *Project) Validate() error {
	if p.PackageName == "" {
		return errors.New("package name is not set")
	}

	empty, err := utils.IsDirEmpty(p.Directory)
	if err != nil {
		return fmt.Errorf("directory empty check failed. dir: %s, err: %w", p.Directory, err)
	}

	if !empty {
		return fmt.Errorf("directory is not empty. dir: %s", p.Directory)
	}

	return nil
}
