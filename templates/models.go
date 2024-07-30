package templates

type Project struct {
	PackageName string

	ProjectDescription string
	ProjectModuleName  string
	ProjectName        string
  Author             string `mapstructure:"author"`

	GoVersion string
}

