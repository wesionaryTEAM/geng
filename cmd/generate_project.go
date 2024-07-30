package cmd

import (
	"github.com/mukezhz/geng/pkg"
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
	pkg.BindStrFlag(projectCmd, "dir", "d", "", "target directory to create project at")
	pkg.BindStrFlag(projectCmd, "goversion", "g", "", "go version for project generation")
}

func createProject(cmd *cobra.Command, args []string) {
	logger := pkg.GetLogger()
	input := pkg.GetConfig[templates.Project]()

	logger.Infof("input: %#v", input)

}
