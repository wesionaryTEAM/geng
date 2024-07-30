package models

import "errors"

type Project struct {
	PackageName string

	ProjectDescription string
	ProjectModuleName  string
	ProjectName        string
	Author             string `mapstructure:"author"`

	GoVersion string `mapstructure:"goversion"`
}

func (p *Project) Validate() error {
	if p.PackageName == "" {
		return errors.New("package name is not set")
	}

	return nil
}
