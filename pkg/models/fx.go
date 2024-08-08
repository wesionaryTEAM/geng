package models

import "fmt"

type Fx struct {
	Directory string `mapstructure:"dir"`
}

func (f *Fx) AutoFill() {
	// TODO: autofill directory if not passed using git root
}

func (f *Fx) Validate() error {
	if f.Directory == "" {
		return fmt.Errorf("directory setting is empty")
	}

	return nil
}
