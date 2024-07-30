package cmd

import (
	"github.com/mukezhz/geng/pkg"
	"github.com/mukezhz/geng/pkg/models"
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

	// TODO: show terminal ui if not present

	logger.Infof("input: %#v", input)

	// validate input data
	if err := input.Validate(); err != nil {
		logger.Fatal("validation for input failed.", "err", err)
	}

	// projectGenerator := pkg.NewProjectGenerator()
	// if err := projectGenerator.Generate(input); err != nil {
	// 	logger.Fatal("project generation failed.", "err", err)
	// }

}
