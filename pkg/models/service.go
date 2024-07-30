package models

type Service struct {
	Directory         string `mapstructure:"dir"`
	ProjectModuleName string `mapstructure:"mod"`

	ServiceType []string `mapstructure:"servicetype"`
}

// Fill fills up some empty data from project name if possible to do so
func (p *Service) AutoFill() {
	if p.Directory == "" {
		p.Directory = "./" + p.ProjectModuleName
	}
}

func (p *Service) Validate() error {
	// TODO: validate input data
	// check directory structures are properly present or not
	// check if mod files are present in the given directory
	// check service choices provided matches from the available in the templates

	return nil
}
