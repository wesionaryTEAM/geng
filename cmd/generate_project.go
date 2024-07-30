package cmd

import (
	"github.com/mukezhz/geng/pkg"
	"github.com/mukezhz/geng/pkg/gen"
	"github.com/mukezhz/geng/pkg/models"
	"github.com/mukezhz/geng/templates"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new project",
	Run:   createProject,
}

func init() {
	pkg.BindStrFlag(projectCmd, "author", "a", "wesionaryTEAM", "author for the project")
	pkg.BindStrFlag(projectCmd, "mod", "m", "", "module name for the project")
	pkg.BindStrFlag(projectCmd, "name", "n", "", "project name")
	pkg.BindStrFlag(projectCmd, "desc", "e", "", "project description")
	pkg.BindStrFlag(projectCmd, "pkg", "p", "", "package name")
}

func createProject(cmd *cobra.Command, args []string) {
	logger := pkg.GetLogger()
	input := pkg.GetConfig[models.Project]()

	// Fill up data items, if can be filled up automatically
	input.AutoFill()

	// TODO: show terminal ui, auto fill up contents already given by flags

	logger.Infof("input: %#v", input)

	// validate input data
	if err := input.Validate(); err != nil {
		logger.Fatal("validation for input failed.", "err", err)
	}

	generator := gen.NewProjectGenerator(input, templates.FS)
	if err := generator.Generate(); err != nil {
		logger.Fatal("project generation failed.", "err", err)
	}

}
