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
	pkg.BindStrFlag(projectCmd, "dir", "d", "", "target directory to create project at")
	pkg.BindStrFlag(projectCmd, "goversion", "g", "", "go version for project generation")
}

func createProject(cmd *cobra.Command, args []string) {
	logger := pkg.GetLogger()
	input := pkg.GetConfig[models.Project]()

	logger.Infof("input: %#v", input)

	// TODO: show terminal ui if not present

	// validate input data
	if err := input.Validate(); err != nil {
		logger.Fatal("validation for input failed.", "err", err)
	}

	// projectGenerator := pkg.NewProjectGenerator()
	// if err := projectGenerator.Generate(input); err != nil {
	// 	logger.Fatal("project generation failed.", "err", err)
	// }

}
