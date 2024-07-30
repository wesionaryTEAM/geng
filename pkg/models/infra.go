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
  // TODO: validate input data
  // check directory structures are properly present or not
  // check if mod files are present in the given directory
  // check infra choices provided matches from the available in the templates

	return nil
}
