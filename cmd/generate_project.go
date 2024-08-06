package cmd

import (
	"github.com/mukezhz/geng/pkg"
	"github.com/mukezhz/geng/pkg/gen"
	"github.com/mukezhz/geng/pkg/models"
	"github.com/mukezhz/geng/pkg/utils"
	"github.com/mukezhz/geng/templates"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new project",
	Run:   createProject,
}

var projectFlags = []utils.FlagItem{
	{Name: "author", Short: "a", Desc: "author for the project"},
	{Name: "mod", Short: "m", Desc: "module name for the project"},
	{Name: "name", Short: "n", Desc: "project name"},
	{Name: "desc", Short: "e", Desc: "project description"},
	{Name: "pkg", Short: "p", Desc: "package name"},
	{Name: "infratype", Short: "i", Desc: "infastructure types (eg; aws, firebase, gcp, gmail, opensearch)", Def: []string{}},
}

func init() {
	utils.SetFlags(projectCmd, projectFlags)
}

func createProject(cmd *cobra.Command, args []string) {
	utils.BindFlag(cmd, projectFlags)

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

	// generate infrastructure after project has been generated
	infraInput := pkg.GetConfig[models.Infrastructure]()
	infraInput.ProjectModuleName = input.ProjectModuleName
	infraInput.Directory = input.Directory

	logger.Infof("infra input: %v", infraInput)

	// validate input data
	if err := infraInput.Validate(); err != nil {
		logger.Fatal("validation for input failed.", "err", err)
	}

	// generate project
	generator := gen.NewProjectGenerator(input, templates.FS)
	if err := generator.Generate(); err != nil {
		logger.Fatal("project generation failed.", "err", err)
	}

	// generate infrastructure
	infraGenerator := gen.NewInfraGenerator(infraInput, templates.FS)
	if err := infraGenerator.Generate(); err != nil {
		logger.Fatal("infrastructure generation failed.", "err", err)
	}

}
