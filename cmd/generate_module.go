package cmd

import (
	"github.com/mukezhz/geng/pkg"
	"github.com/mukezhz/geng/pkg/gen"
	"github.com/mukezhz/geng/pkg/models"
	"github.com/mukezhz/geng/pkg/utils"
	"github.com/mukezhz/geng/templates"
	"github.com/spf13/cobra"
)

var moduleCmd = &cobra.Command{
	Use:   "add:mod",
	Short: "Add a new module in a project. (service, repo, module, controllers, routes etc.)",
	Run:   createModule,
}

var moduleCmdFlags = []utils.FlagItem{
	{Name: "name", Short: "n", Desc: "module name to generate"},
}

func init() {
	utils.SetFlags(moduleCmd, moduleCmdFlags)
}

func createModule(cmd *cobra.Command, args []string) {
	utils.BindFlag(cmd, projectFlags)

	logger := pkg.GetLogger()
	input := pkg.GetConfig[models.Module]()

	if err := input.AutoFill(); err != nil {
		logger.Fatal(err)
	}

	if err := input.Validate(); err != nil {
		logger.Fatal(err)
	}

	generator := gen.NewModuleGenerator(input, templates.FS)
	if err := generator.Generate(); err != nil {
		logger.Fatal("module generation error", "err", err)
	}
}
