package models

type Infrastructure struct {
	Directory         string `mapstructure:"dir"`
	ProjectModuleName string `mapstructure:"mod"`

	InfraType []string `mapstructure:"infratype"`
}

// Fill fills up some empty data from project name if possible to do so
func (p *Infrastructure) AutoFill() {
	if p.Directory == "" {
		p.Directory = "./" + p.ProjectModuleName
	}
}

func (p *Infrastructure) Validate() error {
	// validate input data
	return nil
}
